use std::io;
use std::thread::sleep;
use std::time::Duration;

fn main() {
    let mut healthy = 100;
    let mut hunger = 100;
    let mut penalty = 0;
    println!("healthy={healthy}, hunger={hunger}, penalty={penalty}");
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
        // println!("yorn: {}, unuzluk: {}", yorn, yorn.len()); // debug
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
            hunger = hunger - 1 - penalty;
            println!(
                "Your Stat; \n Healthy:{} \n Hunger:{} \n Penalty:{}",
                healthy, hunger, penalty
            );
            loop {
                let mut Food_selection = String::new();
                io::stdin()
                    .read_line(&mut Food_selection)
                    .expect("Failed to read line");
                let Food_selection = Food_selection.trim();
                if Food_selection == "yes"
                    || Food_selection == "Yes"
                    || Food_selection == "YES"
                    || Food_selection == "y"
                    || Food_selection == "Y"
                {
                    hunger = hunger + 5;
                    break;
                } else if Food_selection == "no"
                    || Food_selection == "No"
                    || Food_selection == "NO"
                    || Food_selection == "n"
                    || Food_selection == "N"
                {
                    hunger = hunger - 5;
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
                    continue;
                }
            }
            if hunger < 0 {
                println!("Your pet is very hungry!");
                healthy -= 1;
                penalty += 1;
                continue;
            } else if hunger < 50 {
                println!("Your pet is hungry!");
                penalty += 1;
                continue;
            } else if hunger > 120 || hunger == 120 {
                println!("Your pet is obese!");
                healthy -= 1;
                penalty += 1;
                continue;
            } else if hunger > 100 || hunger == 100 {
                println!("Your pet is very full!");
                healthy += 1;
                continue;
            } else {
                println!("Your pet is not hungry!");
                continue;
            }
            if healthy == 0 {
                println!("Your pet is dead!");
                break;
            } else if healthy < 50 {
                println!("Your pet is sick!");
                continue;
            } else if healthy < 75 {
                println!("Your pet is fine!");
                continue;
            } else {
                println!("Your pet is very healthy!");
                continue;
            }
        } else {
            println!("hmm what the fuckkk howw!");
            break;
        }
    }
}
