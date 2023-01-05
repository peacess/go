package map_sort_collection

import (
	"testing"

	"github.com/google/btree"
)

const Degree = 2048

func BenchmarkInsert_Btree(b *testing.B) {
	tree := btree.NewG(Degree, btree.Less[int]())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.ReplaceOrInsert(i)
	}
}

func BenchmarkGet_Btree(b *testing.B) {
	tree := btree.NewG(Degree, btree.Less[int]())
	for i := 0; i < b.N; i++ {
		tree.ReplaceOrInsert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tree.Get(i)
	}
}

func BenchmarkDelete_Btree(b *testing.B) {
	tree := btree.NewG(Degree, btree.Less[int]())
	for i := 0; i < b.N; i++ {
		tree.ReplaceOrInsert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tree.Delete(i)
	}
}
