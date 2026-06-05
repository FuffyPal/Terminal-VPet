package main

import "fmt"

func (p *Pet) eat(eat int) {
	if p.hunger > 0 {
		p.hunger -= eat
		if p.hunger < 0 {
			p.hunger = 0
		}
	}
}

func (p *Pet) demage(demage int) {
	if p.healty > 0 {
		p.healty = demage
		if p.healty < 0 {
			p.healty = 0
		}
	}
}

func (p *Pet) heal(heal int) {
	if p.healty > 0 {
		p.healty += heal
		if p.healty > 100 {
			p.healty = 100
		}
	}
}

func (p *Pet) check() {
	if p.healty == 0 {
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
