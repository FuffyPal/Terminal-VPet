package main

import (
	"fmt"
	"os"
)

func (p *Pet) eat(eat int) {
	if p.hunger > minhungry {
		p.hunger -= eat
		if p.hunger < minhungry {
			p.hunger = minhungry
		}
	}
}

func (p *Pet) demage(demage int) {
	if p.healty > minhungry {
		p.healty -= demage
		if p.healty < minhealthy {
			p.healty = minhealthy
		}
	}
}

func (p *Pet) heal(heal int) {
	if p.healty > minhealthy {
		p.healty += heal
		if p.healty > maxhealthy {
			p.healty = maxhealthy
		}
	}
}

func (p *Pet) check() {
	if p.healty == minhealthy {
		p.life = false
		fmt.Println("Pet is die")
		os.Exit(0)
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
