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
