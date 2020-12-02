use std::fs::File;
use std::io::prelude::*;
use std::io::{self, BufReader};

fn main() -> io::Result<()> {
    let file = File::open("input.txt")?;
    let file = BufReader::new(file);

    let mut counter_1 = 0;
    let mut counter_2 = 0;

    for line in file.lines() {
        let unwrapped_line = line.unwrap();
        let items: Vec<&str> = unwrapped_line.split(|c| !char::is_alphanumeric(c)).collect();

        let min: usize = items[0].parse().unwrap();
        let max: usize = items[1].parse().unwrap();
        let c = items[2];
        let c_char = c.chars().next();
        let password = items[4];
        
        let occurrences = password.matches(c).count();
        if occurrences >= min && occurrences <= max {
            counter_1 += 1;
        }

        let contains_first = password.chars().nth(min - 1) == c_char;
        let contains_second = password.chars().nth(max - 1) == c_char;
        if contains_first^contains_second {
            counter_2 += 1;
        }
    }

    println!("{}", counter_1);
    println!("{}", counter_2);

    Ok(())
}
