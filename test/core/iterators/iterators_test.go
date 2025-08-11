package iterators

import (
	"markmind/internal/core/iterators"
	"testing"
)

func Test_Length(t *testing.T) {
	var arr = []int{42, 21, 13, 7, 1}
	iter := iterators.Iter(arr)
	length := iter.Length()
	if length != 5 {
		t.Fatalf("Iterator has an unexpected Length: %d\n", length)
	}
}

func Test_Map(t *testing.T) {
	var arr = []int{42, 21, 13, 7, 1}
	mapped := iterators.Map(arr, func(i, it int) int {
		return it + 1
	})
	expectation := []int{43, 22, 14, 8, 2}
	l := mapped.Length()
	if l != len(expectation) {
		t.Fatalf("Iterator has an unexpected Length: %d\n", l)
	}
	for i, it := range mapped {
		if it != expectation[i] {
			t.Fatalf("Iterator has an unexpected Value (%d) at index (%d)\n", it, i)
		}
	}
}

func Test_MapIter(t *testing.T) {
	var arr = []int{42, 21, 13, 7, 1}
	mapped := iterators.MapIter(iterators.Iter(arr), func(i, it int) int {
		return it + 1
	})
	expectation := []int{43, 22, 14, 8, 2}
	l := mapped.Length()
	if l != len(expectation) {
		t.Fatalf("Iterator has an unexpected Length: %d\n", l)
	}
	for i, it := range mapped {
		if it != expectation[i] {
			t.Fatalf("Iterator has an unexpected Value (%d) at index (%d)\n", it, i)
		}
	}
}

func Test_Filter(t *testing.T) {
	var arr = []int{42, 21, 13, 7, 1}
	filtered := iterators.Iter(arr).Filter(func(i, value int) bool {
		return i%2 == 0
	})
	expectation := []int{42, 13, 1}
	l := len(expectation)
	if l != len(expectation) {
		t.Fatalf("Iterator has an unexpected Length: %d\n", l)
	}
	for i, it := range filtered {
		if it != expectation[i] {
			t.Fatalf("Iterator has an unexpected Value (%d) at index (%d)\n", it, i)
		}
	}
}

func Test_Any(t *testing.T) {
	var arr = []int{42, 21, 13, 7, 1}
	iter := iterators.Iter(arr)
	if iter.Any(func(i, value int) bool { return value == 22 }) {
		t.Fatalf("Unexpected Value in iterator\n")
	}
	if !iter.Any(func(i, value int) bool { return value == 21 }) {
		t.Fatalf("Unexpectedly Value in iterator not found\n")
	}
}
