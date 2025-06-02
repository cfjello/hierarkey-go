package main

import (
	"fmt"

	"github.com/cfjello/hierarkey-go/pkg/hierarkey"
)

func main() {
	hk := hierarkey.NewHierarKey(1, 3)

	// Create a map to store our values
	m := make(map[string]string)

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
	m[hk.JumpToLevel("000")] = "Plant"
	m[hk.NextLevel()] = "Flowering Plant"
	m[hk.NextLevel()] = "Magnoliopsida"
	m[hk.NextLevel()] = "Fabales"
	m[hk.NextLevel()] = "Pae/Bean"
	m[hk.NextLevel()] = "Pisum"
	m[hk.NextLevel()] = "Pea"

	// Iterate over the map and print key-value pairs
	for key, value := range m {
		fmt.Printf("%s: %s\n", key, value)
	}
}
