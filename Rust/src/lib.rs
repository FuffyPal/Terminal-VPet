use godot::prelude::*;

struct Pet {
    name: String,
    healthy: i32,
    hunger: i32,
    penalty: i32,
    starvation: i32,
    obesity: i32,
}

impl Pet {
    fn eat(&mut self, eat: i32) {
        self.hunger += eat;
    }
    fn damage(&mut self, damage: i32) {
        self.healthy -= damage;
    }
    fn heal(&mut self, heal: i32) {
        self.healthy += heal;
    }
    fn fine(&mut self, fine: i32) {
        self.penalty += fine;
    }
    fn weight(&mut self, weight: i32) {
        self.obesity += weight;
    }
    fn weaken(&mut self, weaken: i32) {
        self.starvation -= weaken;
    }
    fn clamp_hunger(&mut self) {
        self.hunger = self.hunger.clamp(0, 100);
    }
    fn clamp_health(&mut self) {
        self.healthy = self.healthy.clamp(0, 100);
    }
    fn status_check(&mut self) {
        self.clamp_health();
        self.clamp_hunger();
        match self.healthy {
            0 => {
                println!("Your pet is dead!");
                panic!("Goodbye!");
            },
            1..=50 => {
                println!("Your pet is sick!");
                // println!("Healthy Stat: {}", self.healthy); // debug
            },
            51..=75 => {
                println!("Your pet is fine!");
                // println!("Healthy Stat: {}", pet.healthy); // debug
            },
            76..=100 => {
                println!("Your pet is very healthy!");
                // println!("Healthy Stat: {}", self.healthy); // debug
            },
            _ => {
                println!("Error: Invalid healthy value!");
            }
        }
        match self.hunger {
            0 => {
                println!("Your pet is very hungry!");
                self.weaken(1);
                self.weight(1);
                // println!("Hunger Stat: {}", pet.hunger); // debug
                // println!("Penalty Stat: {}", pet.penalty); // debug
            },
            1..=50 => {
                println!("Your pet is hungry!");
                self.fine(1);
                // println!("Hunger Stat: {}", pet.hunger); // debug
            },
            51..=75 => {
                println!("Your pet is fine!");
                self.fine(-1);
                // println!("Hunger Stat: {}", pet.hunger); // debug
            },
            76..=100 => {
                println!("Your pet is very healthy!");
                self.fine(-1);
                // println!("Hunger Stat: {}", pet.hunger); // debug
            },
            _ => {
                println!("Error: Invalid hunger value!");
            }
        }
    }
    fn pass_time(&mut self) {
        self.hunger = self.hunger - 1 - self.penalty;
        let damage = self.obesity + self.starvation + self.penalty;
        self.damage(damage);
        self.clamp_health();
        self.clamp_hunger();
        self.status_check();
    }
}

fn main() {
    println!("hello world")
}