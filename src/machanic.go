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

func (p *Pet) damage(damage int) {
	if p.healthy > minhealthy {
		p.healthy -= damage
		if p.healthy < minhealthy {
			p.healthy = minhealthy
		}
	}
}

func (p *Pet) heal(heal int) {
	if p.healthy > minhealthy {
		p.healthy += heal
		if p.healthy > maxhealthy {
			p.healthy = maxhealthy
		}
	}
}

func (p *Pet) check() {
	if p.healthy == minhealthy {
		p.life = false
		fmt.Println("Pet is dead")
		os.Exit(0)
	}
}

func (p *Pet) status() {
	fmt.Println("\n=================================")
	fmt.Println("    🐾 Virtual Baby Status 🐾   ")
	fmt.Println("=================================")
	fmt.Printf(" 🌟 Name:    %s\n", p.name)
	fmt.Printf(" ❤️  Healthy: %d / 100\n", p.healthy)
	fmt.Printf(" 🍖 Hunger:   %d / 100\n", p.hunger)
	fmt.Println("=================================")
}

func (p *Pet) tick(y int) {
	x := y
	p.hunger += x
	if p.hunger >= maxhungry {
		p.hunger = maxhungry
	}
}
