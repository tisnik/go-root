package main

import (
	"testing"

	"github.com/gammazero/deque"
)

func BenchmarkPushFrontInt(b *testing.B) {
	var q deque.Deque[int]

	for i := 0; i < b.N; i++ {
		q.PushFront(i)
	}
}

func BenchmarkPushBackInt(b *testing.B) {
	var q deque.Deque[int]

	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
}

func fillInDequeFromBack(b *testing.B) deque.Deque[int] {
	var q deque.Deque[int]

	b.StopTimer()

	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}

	b.StartTimer()
	return q
}

func fillInDequeFromFront(b *testing.B) deque.Deque[int] {
	var q deque.Deque[int]

	b.StopTimer()

	for i := 0; i < b.N; i++ {
		q.PushFront(i)
	}

	b.StartTimer()
	return q
}

func BenchmarkPopFrontInt(b *testing.B) {
	q := fillInDequeFromBack(b)

	for i := 0; i < b.N; i++ {
		val := q.PopFront()
		if val != i {
			b.Fail()
		}
	}
}

func BenchmarkPopBackInt(b *testing.B) {
	q := fillInDequeFromFront(b)

	for i := 0; i < b.N; i++ {
		val := q.PopBack()
		if val != i {
			b.Fail()
		}
	}
}

func BenchmarkFrontInt(b *testing.B) {
	q := fillInDequeFromBack(b)

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		val := q.Front()
		if val != 0 {
			b.Fail()
		}
	}
}

func BenchmarkBackInt(b *testing.B) {
	q := fillInDequeFromFront(b)

	for i := 0; i < b.N; i++ {
		val := q.Back()
		if val != 0 {
			b.Fail()
		}
	}
}

func BenchmarkRotateInt(b *testing.B) {
	q := fillInDequeFromFront(b)

	for i := 0; i < b.N; i++ {
		q.Rotate(1)
	}
}

func BenchmarkAtInt(b *testing.B) {
	q := fillInDequeFromFront(b)

	for i := 0; i < b.N; i++ {
		// at from the middle
		val := q.At(q.Len() / 2)
		if val < 0 || val > b.N {
			b.Fail()
		}
	}
}

func BenchmarkInsertInt(b *testing.B) {
	q := fillInDequeFromFront(b)

	for i := 0; i < b.N; i++ {
		// insert in the middle
		q.Insert(q.Len()/2, i)
	}
}

func BenchmarkRemoveInt(b *testing.B) {
	q := fillInDequeFromFront(b)

	for i := 0; i < b.N; i++ {
		// remove from the middle
		q.Remove(q.Len() / 2)
	}
}
