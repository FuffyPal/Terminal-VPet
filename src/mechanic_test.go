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
