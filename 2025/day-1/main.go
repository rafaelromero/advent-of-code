package main

import (
	"fmt"
)

type key struct {
	direction string
	clicks    int
}
type Dial struct {
	position int
	distanceFromZero int
	keys     []key
}

func main() {
	fmt.Println("Day 1: Open Safe")

	dial := Dial{position: 20}
	dial.distanceFromZero = dial.position
	dial.keys = loadKey()

	for _, key := range dial.keys {
		dial.position = moveDial(key.direction, key.clicks, &dial.position)
		fmt.Printf("Dial position: %d\n", dial.position)
	}
	//position := moveDial("R", 5, &dial.position)

}

func loadKey() []key {

	key1 := key{direction: "L", clicks: 10}
	key2 := key{direction: "R", clicks: 5}

	keys := []key{key1, key2}

	return keys
}

func moveDial(direction string, clicks int, currentPosition *int) int {
	switch direction {
	case "L":
		(*currentPosition) -= clicks
	case "R":
		(*currentPosition) += clicks

	}

	return *currentPosition

}
