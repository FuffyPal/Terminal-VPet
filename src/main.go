package main

import (
	"flag"
	"fmt"
)

const (
	version    = "0.1.0"
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

func main() {
	nameFlag := flag.String("name", "", "Pet is name")
	statusFlag := flag.Bool("status", false, "Pet is status")
	flag.Parse()
	myPet := Pet{healty: 100, hunger: 0, life: true}
	if *nameFlag != "" {
		myPet.name = *nameFlag
	}

	if *statusFlag {
		myPet.status()
		return
	}

	fmt.Printf("\n✨ %s has successfully hatched! Let the adventure begin...\n", myPet.name)
	myPet.status()
}
