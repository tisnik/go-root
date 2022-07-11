// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů ze devadesáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_93/README.md

package main

import (
	"testing"

	"github.com/daichi-m/go18ds/lists"
	"github.com/daichi-m/go18ds/lists/arraylist"
	"github.com/daichi-m/go18ds/lists/doublylinkedlist"
	"github.com/daichi-m/go18ds/lists/singlylinkedlist"
)

func BenchmarkArrayListInsert0Operation(b *testing.B) {
	l := arraylist.New[int]()
	for i := 0; i < b.N; i++ {
		l.Insert(0, i)
	}
}

func BenchmarkSinglyLinkedListInsert0Operation(b *testing.B) {
	l := singlylinkedlist.New[int]()
	for i := 0; i < b.N; i++ {
		l.Insert(0, i)
	}
}

func BenchmarkDoublyLinkedListInsert0Operation(b *testing.B) {
	l := doublylinkedlist.New[int]()
	for i := 0; i < b.N; i++ {
		l.Insert(0, i)
	}
}

func BenchmarkArrayListInsertMiddleOperation(b *testing.B) {
	l := arraylist.New[int]()
	for i := 0; i < b.N; i++ {
		l.Insert(l.Size()/2, i)
	}
}

func BenchmarkSinglyLinkedListInsertMiddleOperation(b *testing.B) {
	l := singlylinkedlist.New[int]()
	for i := 0; i < b.N; i++ {
		l.Insert(l.Size()/2, i)
	}
}

func BenchmarkDoublyLinkedListInsertMiddleOperation(b *testing.B) {
	l := doublylinkedlist.New[int]()
	for i := 0; i < b.N; i++ {
		l.Insert(l.Size()/2, i)
	}
}

func BenchmarkArrayListAppendOperation(b *testing.B) {
	l := arraylist.New[int]()
	for i := 0; i < b.N; i++ {
		l.Add(i)
	}
}

func BenchmarkSinglyLinkedListAppendOperation(b *testing.B) {
	l := singlylinkedlist.New[int]()
	for i := 0; i < b.N; i++ {
		l.Add(i)
	}
}

func BenchmarkDoublyLinkedListAppendOperation(b *testing.B) {
	l := doublylinkedlist.New[int]()
	for i := 0; i < b.N; i++ {
		l.Add(i)
	}
}

func fillInList(l lists.List[int], n int) {
	for i := 0; i < n; i++ {
		l.Add(i)
	}
}

func BenchmarkArrayListRemoveFirstOperation(b *testing.B) {
	l := arraylist.New[int]()
	fillInList(l, b.N)

	for i := 0; i < b.N; i++ {
		l.Remove(0)
	}
}

func BenchmarkSinglyLinkedListRemoveFirstOperation(b *testing.B) {
	l := singlylinkedlist.New[int]()
	fillInList(l, b.N)

	for i := 0; i < b.N; i++ {
		l.Remove(0)
	}
}

func BenchmarkDoublyLinkedListRemoveFirstOperation(b *testing.B) {
	l := doublylinkedlist.New[int]()
	fillInList(l, b.N)

	for i := 0; i < b.N; i++ {
		l.Remove(0)
	}
}

func BenchmarkArrayListRemoveLastOperation(b *testing.B) {
	l := arraylist.New[int]()
	fillInList(l, b.N)

	for i := 0; i < b.N; i++ {
		l.Remove(l.Size() - 1)
	}
}

func BenchmarkSinglyLinkedListRemoveLastOperation(b *testing.B) {
	l := singlylinkedlist.New[int]()
	fillInList(l, b.N)

	for i := 0; i < b.N; i++ {
		l.Remove(l.Size() - 1)
	}
}

func BenchmarkDoublyLinkedListRemoveLastOperation(b *testing.B) {
	l := doublylinkedlist.New[int]()
	fillInList(l, b.N)

	for i := 0; i < b.N; i++ {
		l.Remove(l.Size() - 1)
	}
}

func BenchmarkArrayListGetOperation(b *testing.B) {
	l := arraylist.New[int]()
	fillInList(l, b.N)

	for i := 0; i < b.N; i++ {
		_, exist := l.Get(i)
		if !exist {
			b.Fail()
		}
	}
}

func BenchmarkSinglyLinkedListGetOperation(b *testing.B) {
	l := singlylinkedlist.New[int]()
	fillInList(l, b.N)

	for i := 0; i < b.N; i++ {
		_, exist := l.Get(i)
		if !exist {
			b.Fail()
		}
	}
}

func BenchmarkDoublyLinkedListGetOperation(b *testing.B) {
	l := doublylinkedlist.New[int]()
	fillInList(l, b.N)

	for i := 0; i < b.N; i++ {
		_, exist := l.Get(i)
		if !exist {
			b.Fail()
		}
	}
}
