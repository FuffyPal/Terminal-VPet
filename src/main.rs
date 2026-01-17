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
        let yorn = yorn.trim();
        // println!("yorn: {}, length: {}", yorn, yorn.len()); // debug
        if yorn == "y" || yorn == "Y" || yorn == "Yes" || yorn == "yes" {
            break;
        } else if yorn == "n" || yorn == "N" || yorn == "No" || yorn == "no" {
            println!("Okey again write the pet name.");
            continue;
        } else {
            println!("Invalid input!");
            continue;
        }
    }
    loop {
        let time = Duration::from_secs(5);
        sleep(time);
        let one_tick = time.as_secs();
        if one_tick > 1 || one_tick == 1 {
            pet.hunger = pet.hunger - 1 - pet.penalty;
            let damage = pet.obesity + pet.starvation + pet.penalty;
            pet.damage(damage);
            if pet.healthy < 0 {
                pet.healthy = 0;
                // println!("{}", pet.healthy); // debug
            } else {
            }
            if pet.healthy == 0 {
                println!("Your pet is dead!");
                panic!("Goodbye!");
            } else if pet.healthy < 50 || pet.healthy == 50 {
                println!("Your pet is sick!");
                // println!("Healthy Stat: {}", pet.healthy); // debug
            } else if pet.healthy < 75 || pet.healthy == 75 {
                println!("Your pet is fine!");
                // println!("Healthy Stat: {}", pet.healthy); // debug
            } else if pet.healthy > 100 || pet.healthy == 100 {
                println!("Your pet is very healthy!");
                // println!("Healthy Stat: {}", pet.healthy); // debug
            } else {
                println!("Health!");
            }
            if pet.hunger < 0 {
                pet.hunger = 0;
                // println!("{}", pet.hunger); // debug
            } else if pet.hunger > 100 {
                pet.hunger = 100;
                // println!("{}", pet.hunger); // debug
            } else {
            }
            println!(
                "Your Stat; \n Healthy:{} \n Hunger:{} \n Penalty:{}",
                pet.healthy, pet.hunger, pet.penalty
            );
            loop {
                let mut Food_selection = String::new();
                io::stdin()
                    .read_line(&mut Food_selection)
                    .expect("Failed to read line");
                let Food_selection = Food_selection.trim();
                println!("Your Pet is Hungery: {}", pet.hunger);
                if Food_selection == "yes"
                    || Food_selection == "Yes"
                    || Food_selection == "YES"
                    || Food_selection == "y"
                    || Food_selection == "Y"
                {
                    pet.eat(5);
                    //println!("{}", pet.hunger); // debug
                    break;
                } else if Food_selection == "no"
                    || Food_selection == "No"
                    || Food_selection == "NO"
                    || Food_selection == "n"
                    || Food_selection == "N"
                {
                    // println!("{}", pet.hunger); // debug
                    break;
                } else if Food_selection == "quit"
                    || Food_selection == "Quit"
                    || Food_selection == "QUIT"
                    || Food_selection == "q"
                    || Food_selection == "Q"
                    || Food_selection == "exit"
                    || Food_selection == "Exit"
                    || Food_selection == "EXIT"
                    || Food_selection == "e"
                    || Food_selection == "E"
                    || Food_selection == "end"
                    || Food_selection == "End"
                    || Food_selection == "END"
                    || Food_selection == "stop"
                    || Food_selection == "Stop"
                    || Food_selection == "STOP"
                    || Food_selection == "s"
                    || Food_selection == "S"
                {
                    panic!("Goodbye!");
                } else {
                    println!("Invalid input!");
                }
            }
            if pet.hunger < 0 {
                pet.hunger = 0;
                // println!("{}", pet.hunger); // debug
            } else if pet.hunger > 100 {
                pet.hunger = 100;
                // println!("{}", pet.hunger); // debug
            } else {
            }
            if pet.hunger == 0 {
                println!("Your pet is very hungry!");
                pet.weaken(1);
                pet.weight(1);
                // println!("Hunger Stat: {}", pet.hunger); // debug
                // println!("Penalty Stat: {}", pet.penalty); // debug
            } else if pet.hunger < 50 || pet.hunger == 50 {
                println!("Your pet is hungry!");
                pet.fine(1);
                // println!("Hunger Stat: {}", pet.hunger); // debug
                // println!("Penalty Stat: {}", pet.penalty); // debug
            } else if pet.hunger == 100 {
                println!("Your pet is very obese!");
                pet.heal(1);
                pet.fine(1);
                pet.weight(1);
                // println!("Hunger Stat: {}", pet.hunger); // debug
                // println!("Penalty Stat: {}", pet.penalty); // debug
            } else if pet.hunger < 75 {
                println!("Your pet is normal!");
                pet.heal(1);
                pet.fine(-1);
                // println!("Hunger Stat: {}", pet.hunger); // debug
                // println!("Penalty Stat: {}", pet.penalty); // debug
            } else if pet.hunger > 75 || pet.hunger == 75 {
                println!("Your pet is normal!");
                pet.heal(1);
                pet.fine(-1);
                // println!("Hunger Stat: {}", pet.hunger); // debug
                // println!("Penalty Stat: {}", pet.penalty); // debug
            } else {
                println!("Hunger");
            }
        } else {
            println!("hmm what the fuckkk howw!");
            break;
        }
    }
}
