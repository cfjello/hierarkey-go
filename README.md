# HierarKey for GO

A package to produces a sortable numbering for nodes in a hierarchy. The keys have the format:

```01.02.02.01: Your hierarchy node entry```

The package should probably not be used for very large hierarchies, but for smaller to medium hierarchies it provides a convenient numbering somewhat similar to what you would find in a book:
```
01: Animal
01.01: Vertebrate
01.01.01: Mammal
01.01.01.01: Carnivore
01.01.01.01.01: Cat
01.01.01.01.01.01: Panthera
01.01.01.01.01.01.01: Tiger
01.01.01.02: Primate
01.01.01.02.01: Great Apes
01.01.01.02.01.01: Pongo
01.01.01.02.01.01.01: Orangutan
01.01.01.02.01.02: Homo
01.01.01.02.01.02.01: Human
02: Plant
02.01: Flowering Plant
02.01.01: Magnoliopsida
02.01.01.01: Fabales
02.01.01.01.01: Pae/Bean
02.01.01.01.01.01: Pisum
02.01.01.01.01.01.01: Pea
```

## Usage 

Now create a new hierarKey instance and generate some keys. In this example the numbering on each level starts with 1 and each keys have 2 digits. This program will produce the output shown here above:

```
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
	m[hk.JumpToLevel("000")] = "Plant"
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

```
You can run the examples from the cmd directory: 

`go run main.go A`

## The HierarKey API

While you have to keep track of how you want to assign keys to each of your entries, the HierarKey instance remembers the current assigned numbering for each of the leafs in the hierarchy, enabling it to automatically provide the next available number when you call one of the API functions, as illustrated by this code example:

`go run main.go C`

```
    hk := hierarkey.NewHierarKey(1, 4)
	fmt.Println("Get the root leaf:")
	fmt.Println(hk.NextLeaf())
	fmt.Println("Go up a few levels:")
	fmt.Println(hk.NextLevel())
	fmt.Println(hk.NextLevel())
	fmt.Println(hk.NextLeaf())
	fmt.Println(hk.NextLevel())
	fmt.Println("Go down a few levels:")
	fmt.Println(hk.PrevLevel())
	fmt.Println(hk.PrevLevel())
	fmt.Println("Jump to an existing level:")
	fmt.Println(hk.JumpToLevel("0001.0001.0002"))
	fmt.Println("Jump to an arbitrary level:")
	fmt.Println(hk.JumpToLevel("7.6.5"))
	fmt.Println(hk.NextLeaf())
	fmt.Println("Go a down 2 levels:")
	fmt.Println(hk.PrevLevel(2))
	fmt.Println("Jump to a level in between:")
	fmt.Println(hk.JumpToLevel("2.1"))
	fmt.Println(hk.NextLeaf())
```

The code produces the following output:

```
Get the root leaf:
0001
Go up a few levels:
0001.0001
0001.0001.0001
0001.0001.0002
0001.0001.0002.0001
Go down a few levels:
0001.0001.0003
0001.0002
Jump to an existing level:
0001.0001.0004
Jump to an arbitrary level:
0007.0006.0005
0007.0006.0006
Go a down 2 levels:
0008
Jump to a level in between:
0002.0001
0002.0002
```
## Fetch and number a hierarchy of Music Genres from Wikipedia

`go run main.go B`

This example will give you the output:

```
01:  Rock 
01.01: Active rock
01.02: Adult album alternative
01.03: Soft rock|Adult-oriented rock
01.04: Afro rock
01.05: Album oriented rock
01.06: Alternative rock
01.06.01: Alternative dance
01.06.02: Britpop
01.06.02.01: Post-Britpop
01.06.03: College rock
01.06.04: Dream pop
01.06.04.01: Shoegaze
01.06.04.01.01: Blackgaze
01.06.05: Grunge
01.06.05.01: Post-grunge
01.06.06: Indie rock
01.06.06.01: Dunedin sound
01.06.06.02: Kindie rock
01.06.06.03: Math rock
01.06.06.04: Midwest emo
01.06.06.05: Post-punk revival
01.06.06.06: Slacker rock
(...)
```
Have a look at the example code in the `examples` directory and for more API information (for now) please check the test file: `mod_test.ts`
