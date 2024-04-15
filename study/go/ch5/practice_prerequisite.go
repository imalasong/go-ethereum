package main

import (
	"fmt"
	"sort"
)

var courcetables map[string][]string = map[string][]string{
	"Chinese": {
		"Write",
		"Alphbate",
	},
	"Math": {
		"Chinese",
		"English",
	},
	"English": {
		"Alphbate",
	},
	"Sport": {
		"Run",
		"Skip",
		"Rap",
		"Chinese",
	},
}

func main() {

	var keys []string
	for key, _ := range courcetables {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var result []string
	seen := make(map[string]bool)

	var deepFirst func([]string)
	deepFirst = func(strings []string) {
		for _, s := range strings {
			if !seen[s] {
				seen[s] = true
				children := courcetables[s]
				if children != nil {
					deepFirst(children)
				}
				result = append(result, s)
			}
		}
	}

	deepFirst(keys)

	fmt.Printf("result:%s\n", result)
}
