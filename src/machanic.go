package main

import "fmt"

func (p *Pet) eat() {
	if p.hungry > 0 {
		p.hungry += 1
	}
}

func (p *Pet) status() {
	fmt.Print("Pet \n", "Pet Name: ", p.name, "\nPet healty: ", p.healty, "\nPet hungry: ", p.hungry, "\n")
}
