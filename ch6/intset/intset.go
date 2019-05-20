package intset

import (
	"bytes"
	"fmt"
)

type intset struct {
	words []uint64
}

// 把集合以64个数为一个集合，每个集合是一个int[64]，
// 通过对其中某一位设置值来标记在64个数中的位置，而如果是第一次出现在64个数范围内就在int64数组中相应位置新建一个
func (s *intset) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}
func (s *intset) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}
func (s *intset) UnionWith(t *intset) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}

	}
}

func (s *intset) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
