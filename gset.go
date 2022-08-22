package gset

type StrSet struct {
	data map[string]bool
}

func NewStrSet() *StrSet {
	return &StrSet{
		data: make(map[string]bool),
	}
}

func (s *StrSet) Has(item string) bool {
	return s.data[item]
}

func (s *StrSet) Add(items ...string) *StrSet {
	for ix := range items {
		s.data[items[ix]] = true
	}
	return s
}

func (s *StrSet) Del(items ...string) *StrSet {
	for ix := range items {
		delete(s.data, items[ix])
	}

	return s
}

func (s *StrSet) Len() int {
	return len(s.data)
}

func (s *StrSet) ToArr() []string {
	var (
		arr = make([]string, len(s.data))
		i   int
	)
	for item := range s.data {
		arr[i] = item
		i++
	}
	return arr
}

type Int64Set struct {
	data map[int64]bool
}

func NewInt64Set() *Int64Set {
	return &Int64Set{
		data: make(map[int64]bool),
	}
}

func (s *Int64Set) Has(item int64) bool {
	return s.data[item]
}

func (s *Int64Set) Add(items ...int64) *Int64Set {
	for ix := range items {
		s.data[items[ix]] = true
	}
	return s
}

func (s *Int64Set) Del(items ...int64) *Int64Set {
	for ix := range items {
		delete(s.data, items[ix])
	}
	return s
}

func (s *Int64Set) Len() int {
	return len(s.data)
}

func (s *Int64Set) ToArr() []int64 {
	var (
		arr = make([]int64, len(s.data))
		i   int
	)
	for item := range s.data {
		arr[i] = item
		i++
	}
	return arr
}
