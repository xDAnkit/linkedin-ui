package main

import "fmt"

func main() {
	dummyMap := map[string]int{
		"one": 3,
	}

	updateMap(dummyMap)

	fmt.Println(dummyMap)
}

func updateMap(m map[string]int) {

	m["five"] = 4
	m["one"] = 8
}