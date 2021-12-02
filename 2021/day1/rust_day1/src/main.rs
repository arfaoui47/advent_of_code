use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
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
    print!("{}", incr);
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}