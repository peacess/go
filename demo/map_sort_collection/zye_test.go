package map_sort_collection

import (
	"testing"

	"github.com/zyedidia/generic"
	"github.com/zyedidia/generic/btree"
)

func BenchmarkInsert_GenericBtree(b *testing.B) {
	tree := btree.New[int, int](generic.Less[int])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(i, i)
	}
}

func BenchmarkGet_GenericBtree(b *testing.B) {
	tree := btree.New[int, int](generic.Less[int])
	for i := 0; i < b.N; i++ {
		tree.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tree.Get(i)
	}
}

func BenchmarkDelete_GenericBtree(b *testing.B) {
	tree := btree.New[int, int](generic.Less[int])
	for i := 0; i < b.N; i++ {
		tree.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Remove(i)
	}
}
