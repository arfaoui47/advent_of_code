use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    part1();
    part2();
}

fn part1() {
    if let Ok(lines) = read_lines("../../input.txt") {
        let mut h = 0;
        let mut v = 0;
        for line in lines {
            if let Ok(instruction) = line {
                let vec = instruction.split(" ").collect::<Vec<&str>>();
                let direction = vec[0];
                let value = vec[1].parse::<i32>().unwrap();
                if direction == "forward" {
                    h += value;
                } else if direction == "up" {
                    v -= value;
                } else if direction == "down" {
                    v += value;
                }
            }
        }
        println!("{}", h * v);
    }

}

fn part2() {
    if let Ok(lines) = read_lines("../../input.txt") {
        let mut h = 0;
        let mut v = 0;
        let mut aim = 0;
        for line in lines {
            if let Ok(instruction) = line {
                let vec = instruction.split(" ").collect::<Vec<&str>>();
                let direction = vec[0];
                let value = vec[1].parse::<i32>().unwrap();
                if direction == "forward" {
                    h += value;
                    v += aim * value;
                } else if direction == "up" {
                    aim -= value;
                } else if direction == "down" {
                    aim += value;
                }
            }
        }
        println!("{}", h * v);
    }
}
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
