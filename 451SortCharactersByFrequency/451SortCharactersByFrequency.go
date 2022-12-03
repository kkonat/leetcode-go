package main

import (
	"bytes"
	"fmt"
)

func frequencySort(s string) string {

	// count frequencies
	fr := map[byte]int{}
	for i := 0; i < len(s); i++ {
		fr[s[i]]++
	}

	// create slice of key,var
	type kv struct {
		byte
		int
	}
	keys := make([]kv, 0, len(fr))
l1:
	for key, f := range fr {
		if len(keys) != 0 {
			for i := 0; i < len(keys); i++ {
				if f > keys[i].int {
					keys = append(keys[:i+1], keys[i:]...)
					keys[i] = kv{key, f}
					continue l1
				}
			}
		}
		keys = append(keys, kv{key, f})
	}

	// gen result
	var ret bytes.Buffer
	for _, k := range keys {
		for l := 0; l < k.int; l++ {
			ret.WriteByte(k.byte)
		}
	}
	return ret.String()
}

func main() {
	fmt.Println(frequencySort("tee"))
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}
