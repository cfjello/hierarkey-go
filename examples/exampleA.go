package examples

import (
	"fmt"
	"sort"

	"github.com/cfjello/hierarkey-go/hierarkey/pkg/hierarkey/pkg/hierarkey"
)

func RunExampleA() {

	// Create a map to store our values
	m := make(map[string]string)
	hk := hierarkey.NewHierarKey(1, 2)
	m[hk.NextLeaf()] = "Animal"
	m[hk.NextLevel()] = "Vertebrate"
	m[hk.NextLevel()] = "Mammal"
	m[hk.NextLevel()] = "Carnivore"
	m[hk.NextLevel()] = "Cat"
	m[hk.NextLevel()] = "Panthera"
	m[hk.NextLevel()] = "Tiger"
	m[hk.PrevLevel(3)] = "Primate"
	m[hk.NextLevel()] = "Great Apes"
	m[hk.NextLevel()] = "Pongo"
	m[hk.NextLevel()] = "Orangutan"
	m[hk.PrevLevel(1)] = "Homo"
	m[hk.NextLevel()] = "Human"
	m[hk.JumpToLevel("01")] = "Plant"
	m[hk.NextLevel()] = "Flowering Plant"
	m[hk.NextLevel()] = "Magnoliopsida"
	m[hk.NextLevel()] = "Fabales"
	m[hk.NextLevel()] = "Pae/Bean"
	m[hk.NextLevel()] = "Pisum"
	m[hk.NextLevel()] = "Pea"

	// Iterate over the map and print key-value pairs in sorted order
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	// Sort the keys
	sort.Strings(keys)

	// Print the map in sorted order
	for _, key := range keys {
		fmt.Printf("%s: %s\n", key, m[key])
	}
}
