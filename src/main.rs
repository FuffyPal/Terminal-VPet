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
        if self.hunger < 0 {
            self.hunger = 0;
            // println!("{}", pet.hunger); // debug
        } else if self.hunger > 100 {
            self.hunger = 100;
            // println!("{}", pet.hunger); // debug
        } else {
        }
    }
    fn clamp_health(&mut self) {
        if self.healthy < 0 {
            self.healthy = 0;
            // println!("{}", self.healthy); // debug
        } else {
        }
    }
    fn status_check(&mut self) {
        self.clamp_health();
        self.clamp_hunger();
        if self.healthy == 0 {
            println!("Your pet is dead!");
            panic!("Goodbye!");
        } else if self.healthy < 50 || self.healthy == 50 {
            println!("Your pet is sick!");
            // println!("Healthy Stat: {}", self.healthy); // debug
        } else if self.healthy < 75 || self.healthy == 75 {
            println!("Your pet is fine!");
            // println!("Healthy Stat: {}", pet.healthy); // debug
        } else if self.healthy > 100 || self.healthy == 100 {
            println!("Your pet is very healthy!");
            // println!("Healthy Stat: {}", self.healthy); // debug
        } else {
            println!("Error: Invalid healthy value!");
        }
        if self.hunger == 0 {
            println!("Your pet is very hungry!");
            self.weaken(1);
            self.weight(1);
            // println!("Hunger Stat: {}", pet.hunger); // debug
            // println!("Penalty Stat: {}", pet.penalty); // debug
        } else if self.hunger < 50 || self.hunger == 50 {
            println!("Your pet is hungry!");
            self.fine(1);
            // println!("Hunger Stat: {}", pet.hunger); // debug
            // println!("Penalty Stat: {}", pet.penalty); // debug
        } else if self.hunger == 100 {
            println!("Your pet is very obese!");
            self.heal(1);
            self.fine(1);
            self.weight(1);
            // println!("Hunger Stat: {}", pet.hunger); // debug
            // println!("Penalty Stat: {}", pet.penalty); // debug
        } else if self.hunger < 75 {
            println!("Your pet is normal!");
            self.heal(1);
            self.fine(-1);
            // println!("Hunger Stat: {}", pet.hunger); // debug
            // println!("Penalty Stat: {}", pet.penalty); // debug
        } else if self.hunger > 75 || self.hunger == 75 {
            println!("Your pet is normal!");
            self.heal(1);
            self.fine(-1);
            // println!("Hunger Stat: {}", pet.hunger); // debug
            // println!("Penalty Stat: {}", pet.penalty); // debug
        } else {
            println!("Error: Invalid hunger level");
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
    // Main Area
    loop {
        let time = Duration::from_secs(5);
        sleep(time);
        let one_tick = time.as_secs();
        if one_tick > 1 || one_tick == 1 {
            pet.pass_time();
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
        } else {
            println!("hmm what the fuckkk howw!");
            break;
        }
    }
}
