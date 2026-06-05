package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	version    = "0.4.0"
	maxhealthy = 100
	minhealthy = 0
	maxhungry  = 100
	minhungry  = 0
)

type Pet struct {
	name   string
	healty int
	hunger int
	life   bool
}

type Food struct {
	name   string
	energy int
}

func main() {
	nameFlag := flag.String("name", "", "Pet is name")
	statusFlag := flag.Bool("status", false, "Pet is status")
	versionFlag := flag.Bool("version", false, "Apps version")
	eatFlag := flag.Int("eat", 0, "pet eating x(1,2,3,4)")
	flag.Parse()

	myPet := Pet{healty: 100, hunger: 0, life: true}
	var foodList = []Food{
		{name: "Omlet 🍳", energy: 20},
		{name: "Fish 🐟", energy: 15},
		{name: "Meat 🥩", energy: 35},
		{name: "Apple 🍎", energy: 5},
	}
	if *eatFlag >= 1 && *eatFlag <= len(foodList) {
		choosfood := foodList[*eatFlag-1]

		fmt.Printf("\nYou fed %s with %s!\n", myPet.name, choosfood.name)

		myPet.eat(choosfood.energy)
	} else if *eatFlag != 0 {
		fmt.Println("Invalid choice! Your pet stayed hungry.")
	}

	if *nameFlag != "" {
		myPet.name = *nameFlag
	}

	if *statusFlag {
		myPet.status()
		return
	}

	if *versionFlag {
		fmt.Printf("vpet version: %s \n", version)
		os.Exit(0)
	}

	fmt.Printf("\n✨ %s has successfully hatched! Let the adventure begin...\n", myPet.name)
}
