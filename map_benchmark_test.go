package ordered_sync_map_test

import (
	"fmt"
	mp "github.com/m-murad/ordered-sync-map"
	"testing"
)

func getPopulatedOrderedSyncMap(size int) *mp.Map {
	m := mp.New()
	populateOrderedSyncMap(m, size)
	return m
}

func populateOrderedSyncMap(m *mp.Map, size int) {
	for i := 0; i < size; i++ {
		m.Put(i, i)
	}
}

func BenchmarkOrderedSyncMapGet(b *testing.B) {
	for n := 2; n <= 2048; n *= 2 {
		m := initMap()
		populateOrderedSyncMap(m, 10)

		b.Run(fmt.Sprintf("Get form ordered_sync_map - %d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = m.Get(4)
			}
		})
	}
}

func BenchmarkOrderedSyncMapPut(b *testing.B) {
	for n := 0; n < 10; n++ {
		m := mp.New()
		b.Run(fmt.Sprintf("Put in ordered_sync_map - %d", n), func(b *testing.B) {
			populateOrderedSyncMap(m, b.N)
		})
	}
}

func BenchmarkOrderedSyncMapDelete(b *testing.B) {
	for n := 1; n < 10; n++ {
		b.Run(fmt.Sprintf("Delete form ordered_sync_map - %d", n), func(b *testing.B) {
			b.StopTimer()
			size := b.N
			m := getPopulatedOrderedSyncMap(size)
			b.StartTimer()
			for i := 0; i < size; i++ {
				m.Delete(i)
			}
		})
	}
}

func BenchmarkOrderedSyncMapUnorderedTraversal(b *testing.B) {
	for n := 1; n < 5; n++ {
		b.Run(fmt.Sprintf("Traverse ordered_sync_map randomly - %d", n), func(b *testing.B) {
			b.StopTimer()
			size := b.N
			m := getPopulatedOrderedSyncMap(size)
			b.StartTimer()
			m.UnorderedRange(func(key interface{}, value interface{}) {})
		})
	}
}

func BenchmarkOrderedSyncMapOrderedTraversal(b *testing.B) {
	for n := 1; n < 5; n++ {
		b.Run(fmt.Sprintf("Traverse ordered_sync_map in order - %d", n), func(b *testing.B) {
			b.StopTimer()
			size := b.N
			m := getPopulatedOrderedSyncMap(size)
			b.StartTimer()
			m.OrderedRange(func(key interface{}, value interface{}) {})
		})
	}
}
