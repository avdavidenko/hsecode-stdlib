package lru

import "container/list"

type Cache struct {
	capacity int
	elems    *list.List
	ht       map[int]*list.Element
}

type pair struct {
	key   int
	value int
}

func New(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		elems:    list.New(),
		ht:       make(map[int](*list.Element)),
	}
}

func (cache *Cache) get(key int) *list.Element {
	el, ok := cache.ht[key]
	if !ok {
		return nil
	}

	cache.elems.MoveToFront(el)
	return el
}

func (cache *Cache) Get(key int) (int, bool) {
	el := cache.get(key)
	if el == nil {
		return 0, false
	}
	return el.Value.(pair).value, true
}

func (cache *Cache) Put(key int, value int) {
	if cache.capacity <= 0 {
		return
	}

	el := cache.get(key)
	if el != nil {
		el.Value = pair{key: key, value: value}
		return
	}

	cache.elems.PushFront(pair{key: key, value: value})
	cache.ht[key] = cache.elems.Front()

	if cache.elems.Len() <= cache.capacity {
		return
	}

	key = cache.elems.Back().Value.(pair).key
	delete(cache.ht, key)
	cache.elems.Remove(cache.elems.Back())
}
