package hashmap

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := make([]string, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, m[key])
	}
}
