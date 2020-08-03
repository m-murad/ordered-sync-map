package ordered_sync_map_test

import (
	mp "github.com/m-murad/ordered-sync-map"
	"testing"
)

func initMap() *mp.Map {
	return mp.New()
}

func TestGetPutDelete(t *testing.T) {
	m := initMap()

	if _, ok := m.Get("some-key"); ok {
		t.Fatal(`Get for key "some-key" should return nil and false`)
	}

	if _, ok := m.Get(123); ok {
		t.Fatal(`Get for key 123 should return nil and false`)
	}

	m.Put("key", 123)
	m.Put(123, "key")

	if _, ok := m.Get("some-key"); ok {
		t.Fatal(`Get for key "some-key" should still return nil and false`)
	}

	if val, ok := m.Get("key"); !ok {
		t.Fatal(`2nd value returned by Get for key "key" should be true`)
	} else if val == nil {
		t.Fatal(`1st value returned by Get for key "key" should be non nil`)
	} else {
		intVal, ok := val.(int)
		if !ok || intVal != 123 {
			t.Fatal(`1st value returned by Get for key "key" should int 123`)
		}
	}

	if val, ok := m.Get(123); !ok {
		t.Fatal("2nd value returned by Get for key 123 should be true")
	} else if val == nil {
		t.Fatal("1st value returned by Get for key 123 should be non nil")
	} else {
		strVal, ok := val.(string)
		if !ok || strVal != "key" {
			t.Fatal(`1st value returned by Get for key "key" should int 123`)
		}
	}

	if exists := m.Delete(123); !exists {
		t.Fatal("Delete for key 123 should return true")
	}

	if exists := m.Delete(123); exists {
		t.Fatal("Delete for key 123 on second time should return false")
	}

	if val, ok := m.Get(123); val != nil || ok {
		t.Fatal("Get for key 123 after calling delete should return nil and false")
	}
}

func TestUnorderedRange(t *testing.T) {
	m := initMap()

	kvs := map[interface{}]interface{}{
		"key":            123,
		123:              "key",
		"some-key":       "val 1",
		"some-other-key": "val_2",
		56.11:            true,
	}

	var insertCount int
	for k, v := range kvs {
		insertCount++
		m.Put(k, v)
	}

	var rangeCount int
	rangeFunc := func(key interface{}, val interface{}) {
		rangeCount++
		if kvs[key] != val {
			t.Fatalf("Value mismatch for key %s. In standard map: %v, In ordered_sync_map %v.", key, kvs[key], val)
		}
	}

	m.UnorderedRange(rangeFunc)

	if insertCount != rangeCount {
		t.Fatalf("Range count mismatch. Expected %d got %d", insertCount, rangeCount)
	}
}

func TestOrderedRange(t *testing.T) {
	kvs := [][]interface{}{
		{"key", 123, "some-key", "some-other-key", 56.11}, //keys
		{123, "key", "val 1", "val_2", true},              //values
	}

	m := mp.New()
	for i, _ := range kvs[0] {
		m.Put(kvs[0][i], kvs[1][i])
	}

	var rangeCount int
	rangeFunc := func(key interface{}, val interface{}) {
		if kvs[0][rangeCount] != key {
			t.Fatalf("Key sequesnce mismatic at position %d. Extected %v, received %v.", rangeCount+1, kvs[0][rangeCount], key)
		}
		if kvs[1][rangeCount] != val {
			t.Fatalf("Value sequesnce mismatic at position %d. Extected %v, received %v.", rangeCount+1, kvs[1][rangeCount], val)
		}
		rangeCount++
	}
	m.OrderedRange(rangeFunc)
}
