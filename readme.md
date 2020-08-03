## ordered-sync-map

ordered-sync-map is a package that implements a goroutine-safe ordered map.

### usage

```
import (
    "log"
    mp "github.com/m-murad/ordered-sync-map"
)

var m *mp.Map

func main() {
    // Initialise a new thread-safe ordered Map.
    m = mp.New()
    
    // Insert element in the Map.
    m.Put(key, value)

    // Get will retrive the value for a key in the Map.
    val, ok := m.Get(key)

    // Delete will delete an entry form the Map.
    existed := m.Delete(key)

    // UnorderedRange will iterate over the Map in a random sequence.
    // This is same as ranging over a map using the "for range" syntax.
    m.UnorderedRange(func(key interface{}, val interface{}) {
        log.Println("Key - %v, Value - %v", key, val)
    })

    // OrderedRange will iterate over the Map in the sequence in which 
    // elements were added.
    // This is probable slower than UnorderedRange.
    m.OrderedRange(func(key interface{}, val interface{}) {
        log.Println("Key - %v, Value - %v", key, val)
    })
}
```
