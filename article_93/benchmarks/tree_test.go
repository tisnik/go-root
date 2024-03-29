// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devadesátá třetí část
//    Knihovny s implementací generických datových typů pro programovací jazyk Go (2)
//    https://www.root.cz/clanky/knihovny-s-implementaci-generickych-datovych-typu-pro-programovaci-jazyk-go-2/
//
// Seznam příkladů z devadesáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_93/README.md

package main

import (
	"testing"

	"github.com/daichi-m/go18ds/trees/avltree"
	"github.com/daichi-m/go18ds/trees/btree"
	"github.com/daichi-m/go18ds/trees/redblacktree"
)

func BenchmarkAVLTreePutOperation(b *testing.B) {
	t := avltree.NewWithIntComparator[int]()
	for i := 0; i < b.N; i++ {
		t.Put(i, i)
	}
}

func BenchmarkRBTreePutOperation(b *testing.B) {
	t := redblacktree.NewWithIntComparator[int]()
	for i := 0; i < b.N; i++ {
		t.Put(i, i)
	}
}

func BenchmarkBTree3PutOperation(b *testing.B) {
	t := btree.NewWithIntComparator[int](3)
	for i := 0; i < b.N; i++ {
		t.Put(i, i)
	}
}

func BenchmarkBTree4PutOperation(b *testing.B) {
	t := btree.NewWithIntComparator[int](4)
	for i := 0; i < b.N; i++ {
		t.Put(i, i)
	}
}

func BenchmarkBTree5PutOperation(b *testing.B) {
	t := btree.NewWithIntComparator[int](5)
	for i := 0; i < b.N; i++ {
		t.Put(i, i)
	}
}

func BenchmarkBTree9PutOperation(b *testing.B) {
	t := btree.NewWithIntComparator[int](9)
	for i := 0; i < b.N; i++ {
		t.Put(i, i)
	}
}

func BenchmarkBTree99PutOperation(b *testing.B) {
	t := btree.NewWithIntComparator[int](99)
	for i := 0; i < b.N; i++ {
		t.Put(i, i)
	}
}

const MAX = 10000

func BenchmarkAVLTreeGet(b *testing.B) {
	t := avltree.NewWithIntComparator[int]()
	for i := 0; i < MAX; i++ {
		t.Put(i, i)
	}

	for i := 0; i < b.N; i++ {
		_, _ = t.Get(i % MAX)
	}
}

func BenchmarkRBTreeGet(b *testing.B) {
	t := redblacktree.NewWithIntComparator[int]()
	for i := 0; i < MAX; i++ {
		t.Put(i, i)
	}

	for i := 0; i < b.N; i++ {
		_, _ = t.Get(i % MAX)
	}
}

func BenchmarkBTree3Get(b *testing.B) {
	t := btree.NewWithIntComparator[int](3)
	for i := 0; i < MAX; i++ {
		t.Put(i, i)
	}

	for i := 0; i < b.N; i++ {
		_, _ = t.Get(i % MAX)
	}
}

func BenchmarkBTree4Get(b *testing.B) {
	t := btree.NewWithIntComparator[int](4)
	for i := 0; i < MAX; i++ {
		t.Put(i, i)
	}

	for i := 0; i < b.N; i++ {
		_, _ = t.Get(i % MAX)
	}
}

func BenchmarkBTree9Get(b *testing.B) {
	t := btree.NewWithIntComparator[int](9)
	for i := 0; i < MAX; i++ {
		t.Put(i, i)
	}

	for i := 0; i < b.N; i++ {
		_, _ = t.Get(i % MAX)
	}
}

func BenchmarkBTree99Get(b *testing.B) {
	t := btree.NewWithIntComparator[int](99)
	for i := 0; i < MAX; i++ {
		t.Put(i, i)
	}

	for i := 0; i < b.N; i++ {
		_, _ = t.Get(i % MAX)
	}
}
