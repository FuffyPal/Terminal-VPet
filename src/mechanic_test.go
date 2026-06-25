package main

import (
	"testing"
	"time"
)

func TestEat(t *testing.T) {

	my := &Pet{
		Life:   true,
		Hunger: 50,
	}
	my.eat(20)
	expect := 30

	if my.Hunger != expect {
		t.Errorf("The eat() function didn't reduce hunger correctly! Expected: %d, Actual: %d", expect, my.Hunger)
	}
}

func TestFoodEat(t *testing.T) {
	for i := 8; i > 0; i-- {
		c := i
		f := &Food{
			name:   "",
			energy: 0,
			poison: 0,
		}
		my := &Pet{
			Life:   true,
			Hunger: 600,
		}
		var foodList = []Food{
			{name: "Omlet 🍳", energy: 20, poison: 0},
			{name: "Fish 🐟", energy: 15, poison: 0},
			{name: "Meat 🥩", energy: 35, poison: 0},
			{name: "Apple 🍎", energy: 5, poison: 0},
			{name: "Spaghetti 🍝", energy: 30, poison: 0},
			{name: "Pizza 🍕", energy: 40, poison: 0},
			{name: "Hamburger 🍔", energy: 20, poison: 0},
			{name: "Chips 🍟", energy: 10, poison: 0},
		}

		choosfood := foodList[c-1]
		t.Logf("\nYou fed %s with %s!\n", f.name, choosfood.name)
		my.eat(choosfood.energy)

		var expectlist = []struct {
			expecte int
		}{
			{expecte: 580}, // c = 1 (Omlet)
			{expecte: 585}, // c = 2 (Fish)
			{expecte: 565}, // c = 3 (Meat)
			{expecte: 595}, // c = 4 (Apple)
			{expecte: 570}, // c = 5 (Spaghetti)
			{expecte: 560}, // c = 6 (Pizza)
			{expecte: 580}, // c = 7 (Hamburger)
			{expecte: 590}, // c = 8 (Chips)
		}

		var choosexpect struct {
			expecte int
		}

		choosexpect = expectlist[c-1]

		if my.Hunger != choosexpect.expecte {
			t.Errorf("The eat() function didn't reduce hunger correctly! Expected: %d, Actual: %d", choosexpect.expecte, my.Hunger)
		}
	}
}

func TestDamage(t *testing.T) {
	my := &Pet{
		Life:    true,
		Hunger:  80,
		Healthy: 100,
	}
	my.damage(10)
	expect := 90

	if my.Healthy != expect {
		t.Errorf("The demage() func didnt reduce Healthy correctly! Expected: %d, Actual: %d", expect, my.Healthy)
	}
}

func TestRealtime(t *testing.T) {
	my := &Pet{
		Time: "",
	}
	my.CurrentTime()
	expect := time.Now().Format("2006-01-02 15:04:05")
	if my.Time != expect {
		t.Errorf("the Current_Time() func didnt realtime correctly! Expected: %s, Actual: %s", expect, my.Time)
	}
}

func TestCalculationTime(t *testing.T) {
	pastTime := time.Now().Add(-23 * time.Minute)
	my := &Pet{
		Time:   pastTime.Format("2006-01-02 15:04:05"),
		Hunger: 80,
	}

	my.TickCalculation()
	expect := 80 + (4 * 5)
	if my.Hunger != expect {
		t.Errorf("the TickCalculation() func didnt realtime correctly! Expected: %d, Actual: %d", expect, my.Hunger)
	}
}

func TestHeal(t *testing.T) {
	my := &Pet{
		Life:    true,
		Hunger:  20,
		Healthy: 90,
	}
	my.heal(10)
	expect := 100
	if my.Healthy != expect {
		t.Errorf("the heal() didnt increase Healthy correctly! Expected: %d, Actual: %d", expect, my.Healthy)
	}
}
