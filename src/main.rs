use std::io;
use std::thread::sleep;
use std::time::Duration;

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
                self.fine(1);
                // println!("Hunger Stat: {}", pet.hunger); // debug
            },
            76..=100 => {
                println!("Your pet is very healthy!");
                self.fine(1);
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
    let mut pet = Pet {
        name: String::new(),
        healthy: 100,
        hunger: 100,
        penalty: 0,
        starvation: 0,
        obesity: 0,
    };
    println!(
        "healthy={}, \n hunger={}, \n penalty={}",
        pet.healthy, pet.hunger, pet.penalty
    );
    // Pet Name Area
    loop {
        println!("Enter pet name: ");
        let mut name = String::new();
        io::stdin()
            .read_line(&mut name)
            .expect("Failed to read line");
        println!("New Pet name: {name}");
        println!("y/n enter: ");
        let mut yorn = String::new();
        io::stdin()
            .read_line(&mut yorn)
            .expect("Failed to read line");
        match yorn.trim().to_lowercase().as_str() {
            "y" | "yes" => break,
            "n" | "no" => {
                println!("Okey again write the pet name.");
                continue;
            },
            _ => {
                println!("Invalid input!");
                continue;
            }
        };
    }
    // Main Area
    loop {
        let time = Duration::from_secs(5);
        sleep(time);
        let one_tick = time.as_secs();
        if one_tick >= 1 {
            pet.pass_time();
            println!(
                "Your Stat; \n Healthy:{} \n Hunger:{} \n Penalty:{}",
                pet.healthy, pet.hunger, pet.penalty
            );
            loop {
                let mut food_selection = String::new();
                io::stdin()
                    .read_line(&mut food_selection)
                    .expect("Failed to read line");
                println!("Your Pet is Hungery: {}", pet.hunger);
                match food_selection.trim().to_lowercase().as_str() {
                    "y" | "yes" => {
                        pet.eat(5);
                        //println!("{}", pet.hunger); // debug
                        break;
                    }
                    "n" | "no" => {
                        println!("Okey again write the pet name.");
                        continue;
                    },
                    "q" | "quit" | "e" | "exit" => {
                        panic!("Goodbye!");
                    },
                    _ => {
                        println!("Invalid input. Please enter 'yes' or 'no'.");
                        continue;
                    }
                }
            }
        } else {
            println!("hmm what the fuckkk howw!");
            break;
        }
    }
}
