package gset

import "testing"

func TestStrSet(t *testing.T) {
	s := NewStrSet()
	vals := s.Add("hello", "world", "steve").Del("steve", "hello").ToArr()
	t.Logf("Values: %v", vals)
}

func TestInt64Str(t *testing.T) {
	s := NewInt64Set()
	vals := s.Add(2022, 8, 10).Del(8).ToArr()
	t.Logf("Values: %v", vals)
}
