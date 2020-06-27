package fulltext

import "strings"

type Index struct {
	idx        map[string]([]int)
	docsNumber int
}

func New(docs []string) *Index {
	idx := new(Index)
	idx.idx = make(map[string]([]int))
	idx.docsNumber = len(docs)
	for i := 0; i < idx.docsNumber; i++ {
		idx.addDoc(i, docs[i])
	}

	return idx
}

func (idx *Index) addDoc(number int, doc string) {
	words := strings.Split(doc, " ")
	for i := 0; i < len(words); i++ {
		docIdxes, ok := idx.idx[words[i]]
		if ok {
			if docIdxes[len(docIdxes)-1] != number {
				idx.idx[words[i]] = append(docIdxes, number)
			}
		} else {
			idx.idx[words[i]] = []int{number}
		}
	}
}

func (idx *Index) Search(query string) []int {
	if len(query) == 0 {
		return []int{}
	}

	words := strings.Split(query, " ")
	if len(words) == 0 {
		return []int{}
	}

	meets := make([]int, idx.docsNumber)

	for i := 0; i < len(words); i++ {
		docIdxes, ok := idx.idx[words[i]]
		if ok {
			for j := 0; j < len(docIdxes); j++ {
				meets[docIdxes[j]]++
			}
		}
	}

	result := make([]int, 0)
	for i := 0; i < idx.docsNumber; i++ {
		if meets[i] == len(words) {
			result = append(result, i)
		}
	}

	return result
}
