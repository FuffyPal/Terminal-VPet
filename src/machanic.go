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
	if p.Hunger > 70 {
		p.Healthy -= damage
		if p.Healthy < minhealthy {
			p.Healthy = minhealthy
		}
	}
}

func (p *Pet) heal(heal int) {
	if p.Hunger < 30 {
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
	p.check()
	if p.Life == true {
		fmt.Println("\n=================================")
		fmt.Println("    🐾 Virtual Baby Status 🐾   ")
		fmt.Println("=================================")
		fmt.Printf(" 🌟 Name:    %s\n", p.Name)
		fmt.Printf(" ❤️  Healthy: %d / 100\n", p.Healthy)
		fmt.Printf(" 🍖 Hunger:   %d / 100\n", p.Hunger)
		fmt.Println("=================================")
	} else {
		fmt.Println("\n=================================")
		fmt.Println("    ✨ 🐾 Rainbow Bridge 🐾 ✨   ")
		fmt.Println("=================================")
		fmt.Printf("  🌈 %s is on the clouds now...\n", p.Name)
		fmt.Println("  💤 Sleeping peacefully in heaven. ")
		fmt.Println("  🍖 Tummy is full, no more pain. ")
		fmt.Println("=================================")
		fmt.Println("  ✨ Try again for a new adventure! ✨")
	}
}

func (p *Pet) tick(y int) {
	x := y
	p.Hunger += x
	if p.Hunger >= maxhungry {
		p.Hunger = maxhungry
	}
	p.check()
	p.damage(10)
	p.heal(9)
}
