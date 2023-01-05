package map_sort_collection

import (
	"testing"

	gbtree "github.com/google/btree"
	"github.com/tidwall/btree"
)

// const Degree = 1024

func BenchmarkInsert_TidwallBtree(b *testing.B) {
	tree := btree.NewBTreeG(gbtree.Less[int]())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Set(i)
	}
}

func BenchmarkGet_TidwallBtree(b *testing.B) {
	tree := btree.NewBTreeG(gbtree.Less[int]())
	for i := 0; i < b.N; i++ {
		tree.Set(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tree.Get(i)
	}
}

func BenchmarkDelete_TidwallBtree(b *testing.B) {
	tree := btree.NewBTreeG(gbtree.Less[int]())
	for i := 0; i < b.N; i++ {
		tree.Set(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tree.Delete(i)
	}
}
