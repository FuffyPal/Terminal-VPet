package main

import (
	"fmt"
	"os"
)

func (p *Pet) eat(eat int) {
	if p.Hunger > minhungry {
		p.Hunger -= eat
		if p.Hunger < minhungry {
			p.Hunger = minhungry
		}
	}
}

func (p *Pet) damage(damage int) {
	if p.Healthy > minhealthy {
		p.Healthy -= damage
		if p.Healthy < minhealthy {
			p.Healthy = minhealthy
		}
	}
}

func (p *Pet) heal(heal int) {
	if p.Healthy > minhealthy {
		p.Healthy += heal
		if p.Healthy > maxhealthy {
			p.Healthy = maxhealthy
		}
	}
}

func (p *Pet) check() {
	if p.Healthy == minhealthy {
		p.Life = false
		fmt.Println("Pet is dead")
		os.Exit(0)
	}
}

func (p *Pet) status() {
	fmt.Println("\n=================================")
	fmt.Println("    🐾 Virtual Baby Status 🐾   ")
	fmt.Println("=================================")
	fmt.Printf(" 🌟 Name:    %s\n", p.Name)
	fmt.Printf(" ❤️  Healthy: %d / 100\n", p.Healthy)
	fmt.Printf(" 🍖 Hunger:   %d / 100\n", p.Hunger)
	fmt.Println("=================================")
}

func (p *Pet) tick(y int) {
	x := y
	p.Hunger += x
	if p.Hunger >= maxhungry {
		p.Hunger = maxhungry
	}
	p.check()
}
