package main

import (
	"fmt"
	"time"
)

const (
	maxhealthy = 100
	minhealthy = 0
	maxhungry  = 100
	minhungry  = 0
)

func (p *Pet) eat(eat int) {
	if p.Life == true {
		if p.Hunger > minhungry {
			p.Hunger -= eat
			if p.Hunger < minhungry {
				p.Hunger = minhungry
			}
		}
	}
}

func (p *Pet) damage(damage int) {
	if p.Life == true {
		if p.Hunger > 70 {
			p.Healthy -= damage
			if p.Healthy < minhealthy {
				p.Healthy = minhealthy
			}
		}
	}
}

func (p *Pet) heal(heal int) {
	if p.Life == true {
		if p.Hunger < 30 {
			p.Healthy += heal
			if p.Healthy > maxhealthy {
				p.Healthy = maxhealthy
			}
		}
	}
}

func (p *Pet) check() {
	if p.Healthy == minhealthy {
		p.Life = false
	}
	if p.Healthy > maxhealthy {
		p.Healthy = maxhealthy
	}
	if p.Hunger > maxhungry {
		p.Hunger = maxhungry
	}
	if p.Healthy < minhealthy {
		p.Healthy = minhealthy
	}
	if p.Hunger < minhungry {
		p.Hunger = minhungry
	}
}

func (p *Pet) CurrentTime() {
	p.Current_Time = time.Now().Format("2006-01-02 15:04:05")
}

func (p *Pet) status() {
	if p.Life == true {
		fmt.Println("\n=================================")
		fmt.Println("    🐾 Virtual Baby Status 🐾   ")
		fmt.Println("=================================")
		fmt.Printf(" 🌟 Name:    %s\n", p.Name)
		fmt.Printf(" ❤️  Healthy: %d / 100\n", p.Healthy)
		fmt.Printf(" 🍖 Hunger:   %d / 100\n", p.Hunger)
		fmt.Println("=================================")
	} else {
		fmt.Println("\n========================================")
		fmt.Println("    ✨ 🐾 Rainbow Bridge 🐾 ✨   ")
		fmt.Println("========================================")
		fmt.Printf("  🌈 %s is on the clouds now...\n", p.Name)
		fmt.Println("  💤 Sleeping peacefully in heaven. ")
		fmt.Println("  🍖 Tummy is full, no more pain. ")
		fmt.Println("========================================")
		fmt.Println("  ✨ Try again for a new adventure! ✨")
	}
}

func (p *Pet) tick(y int) {
	p.Hunger += y
	if p.Hunger >= maxhungry {
		p.Hunger = maxhungry
	}
	p.damage(10)
	p.heal(5)
	p.check()
}
