package main

import "fmt"

// const (
// 	version    = "0.1.0"
// 	debug      = true
// 	maxhealthy = 100
// 	minhealthy = 0
// 	maxhungry  = 100
// 	minhungry  = 0
// )

type Pet struct {
	name   string
	healty int
	hunger int
	life   bool
}

func main() {
	myPet := Pet{healty: 100, hunger: 0, life: true}
	fmt.Scanln(&myPet.name)
	myPet.status()
}
