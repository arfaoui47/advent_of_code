use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn part1() {
    // let mut values: Vec<&str>;
    let mut fishes: Vec<i32> = Vec::new();

    if let Ok(lines) = read_lines("../../input.txt") {
        for line in lines {
            if let Ok(instruction) = line {
                let values = instruction.split(",").collect::<Vec<&str>>();

                for i in 0..values.len() {
                    fishes.push(values[i].parse::<i32>().unwrap());
                }
                let mut append = 0;
                for n in 0..18 {
                    println!("day {}: {:?}, {}", n, fishes ,fishes.len());
                    for fish in fishes.iter_mut() {
                        if *fish > 0 {
                            *fish -= 1;
                        }
                        else{
                            *fish = 6;
                            append += 1;
                        }
                    }
                    for _i in 0..append {
                        fishes.push(8);
                    }
                    append = 0;
                }
            }
        }
    }
}

fn main() {
    part1();
}
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
