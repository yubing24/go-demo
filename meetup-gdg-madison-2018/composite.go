package main

import "fmt"

func main() {
	// Instantiate array with literals
	myArray := [4]int{1, 2, 3, 4}

	// there is no while loop in Go
	for i := 0; i < len(myArray); i++ {
		fmt.Printf("%d\n", myArray[i])
	}

	// map
	menu := make(map[string]int) // Initialize the map.
	menu["mocha"] = 4
	menu["pumpkin spice latte"] = 5
	menu["red tea"] = 4
	menu["hot chocolate"] = 3

	fmt.Println(menu)
	delete(menu, "pumpkin spice latte") // Delete from a map by referring to its key.
	// _, isPresent := itsamap["d"] // "_" refers to the first variable,  which we are going to ignore.
	// Note the two return values for the map. The second value returns if the key is present in the map.

	for k, v := range menu {
		fmt.Printf("Key = [%s], Val = [%d]\n", k, v)
	}
}
