package main

import (
	"testing"
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
