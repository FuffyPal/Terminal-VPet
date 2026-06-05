package main

import "fmt"

func (p *Pet) eat() {
	if p.hunger > 0 {
		p.hunger -= 1
	}
}

// func (p *Pet) demage() {

// }

// func (p *Pet) heal() {

// }

func (p *Pet) check() {
	if p.healty > 0 {
		p.life = false
	}
}

func (p *Pet) status() {
	fmt.Println("\n=================================")
	fmt.Println("    🐾 Virtual Baby Status 🐾   ")
	fmt.Println("=================================")
	fmt.Printf(" 🌟 Name:    %s\n", p.name)
	fmt.Printf(" ❤️  Healty:  %d / 100\n", p.healty)
	fmt.Printf(" 🍖 Hunger:   %d / 100\n", p.hunger)
	fmt.Println("=================================")
}
