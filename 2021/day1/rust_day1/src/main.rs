use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    part1();
    part2();
}

fn part1() {
    // File hosts must exist in current path before this produces output
    let mut v = Vec::new();
    if let Ok(lines) = read_lines("../../input.txt") {
        // Consumes the iterator, returns an (Optional) String
        for line in lines {
            if let Ok(number) = line {
                v.push(number.parse::<i32>().unwrap());
            }
        }
    }
    let mut prev = v[0];
    let mut incr = 0;
    for i in &v {
        if *i > prev {
            incr += 1;
        }
        prev = *i;
    }
    println!("{}", incr);
}

fn part2() {
    // File hosts must exist in current path before this produces output
    let mut v = Vec::new();
    if let Ok(lines) = read_lines("../../input.txt") {
        // Consumes the iterator, returns an (Optional) String
        for line in lines {
            if let Ok(number) = line {
                v.push(number.parse::<i32>().unwrap());
            }
        }
    }
    let mut prev = v[0] + v[1] + v[2];
    let mut incr = 0;
    let mut current;
    for i in 1..v.len() - 2 {
        current = v[i] + v[i + 1] + v[i + 2];
        if current > prev {
            incr += 1;
        }
        prev = current;
    }
    print!("{}", incr);
}
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
