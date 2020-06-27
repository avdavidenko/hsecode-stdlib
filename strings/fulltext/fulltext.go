package fulltext

import "sort"

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

func getWord(idx int, doc string) (string, int) {
	for idx < len(doc) && doc[idx] == ' ' {
		idx++
	}

	startIdx := idx
	for idx < len(doc) && doc[idx] != ' ' {
		idx++
	}

	return doc[startIdx:idx], idx
}

func (idx *Index) addDoc(number int, doc string) {
	i := 0
	word := ""
	for {
		word, i = getWord(i, doc)
		if len(word) == 0 {
			break
		}

		docIdxes, ok := idx.idx[word]
		if ok {
			if docIdxes[len(docIdxes)-1] != number {
				idx.idx[word] = append(docIdxes, number)
			}
		} else {
			idx.idx[word] = []int{number}
		}
	}
}

func (idx *Index) Search(query string) []int {
	if len(query) == 0 {
		return []int{}
	}

	result := make([]int, 0)
	for n := 0; n < idx.docsNumber; n++ {
		i := 0
		word := ""
		count := 0
		for {
			word, i = getWord(i, query)
			if len(word) == 0 {
				if count > 0 {
					result = append(result, n)
				}
				break
			}

			count++
			docIdxes, ok := idx.idx[word]
			if !ok {
				break
			}

			pos := sort.SearchInts(docIdxes, n)
			if pos >= len(docIdxes) || docIdxes[pos] != n {
				break
			}
		}
	}

	return result
}
