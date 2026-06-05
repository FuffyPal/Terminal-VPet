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
	hungry int
}

func main() {
	fmt.Println("Hello, World")
	myPet := Pet{name: "bal", healty: 100, hungry: 100}
	myPet.status()
}
