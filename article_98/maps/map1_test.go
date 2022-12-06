package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

type UUID string

func BenchmarkInsertIntoPreallocatedMapUUIDKey(b *testing.B) {
	m := make(map[UUID]time.Time, b.N)
	t := time.Now()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		id := UUID(uuid.New().String())
		b.StartTimer()
		m[id] = t
	}
}

func BenchmarkInsertIntoEmptyMapUUIDKey(b *testing.B) {
	m := map[UUID]time.Time{}
	t := time.Now()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		id := UUID(uuid.New().String())
		b.StartTimer()
		m[id] = t
	}
}
