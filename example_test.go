package ordered_sync_map_test

import (
	"fmt"

	ordered_sync_map "github.com/m-murad/ordered-sync-map"
)

func ExampleNew() {

	mp := ordered_sync_map.New[string, string]()

	mp.Put("k1", "v1")

	v, ok := mp.Get("k1")
	fmt.Println(v, ok)

	ok = mp.Delete("k2")
	fmt.Println(ok)

	mp.UnorderedRange(func(key, value string) {
		fmt.Println(key, value)
	})

	mp.OrderedRange(func(key, value string) {
		fmt.Println(key, value)
	})

	len := mp.Length()
	fmt.Println(len)

	v, ok = mp.GetOrPut("k1", "v2")
	fmt.Println(v, ok)

	v, ok = mp.GetAndDelete("k1")
	fmt.Println(v, ok)
}
