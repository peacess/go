package map_sort_collection

import (
	"testing"

	"github.com/wzshiming/llrb"
)

func BenchmarkInsert_LLRB(b *testing.B) {
	tree := llrb.NewTree[int, int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(i, i)
	}
}

func BenchmarkGet_LLRB(b *testing.B) {
	tree := llrb.NewTree[int, int]()
	for i := 0; i < b.N; i++ {
		tree.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tree.Get(i)
	}
}

func BenchmarkDelete_LLRB(b *testing.B) {
	tree := llrb.NewTree[int, int]()
	for i := 0; i < b.N; i++ {
		tree.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tree.Delete(i)
	}
}
