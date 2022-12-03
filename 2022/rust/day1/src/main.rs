use std::fs::File;
use std::io::{self, BufRead};
use std::cmp;

fn main() {
    let result = part1();
    match result {
        Ok(s) => println!("Part 1 {}\n", s),
        Err(err) => panic!("Failure {}\n", err)
    }

    let result2 = part2();
    match result2 {
        Ok(s) => println!("part 2 {}\n", s),
        Err(err) => panic!("Failure {}\n", err)
    }
}

fn part1() -> Result<u32, Box<dyn std::error::Error>> {
    let file = File::open("./day1input.txt")?;
    let reader = io::BufReader::new(file);
    let mut current = 0;
    let mut result: u32 = 0;

    for line in reader.lines() {

        match line {
            Ok(s) => {
                match s.parse::<u32>() {
                    Ok(val) => {
                        current = current + val
                    },
                    Err(_e) => {
                        result = cmp::max(result, current);
                        current = 0;
                    }
                }
            }
            Err(err) => panic!("{}", err)
        }
    }

    Ok(result)
}

fn part2() -> Result<u32, Box<dyn std::error::Error>> {
    let file = File::open("./day1input.txt")?;
    let reader = io::BufReader::new(file);
    let mut current = 0;
    let mut results: Vec<u32> = vec![];

    for line in reader.lines() {

        match line {
            Ok(s) => {
                match s.parse::<u32>() {
                    Ok(val) => {
                        current = current + val
                    },
                    Err(_e) => {
                        results.push(current);
                        current = 0;
                    }
                }
            }
            Err(err) => panic!("{}", err)
        }
    }

    results.sort_unstable();

    // In real life we should verify there are at least 3 elements
    Ok(results[results.len() - 1] + results[results.len() - 2] + results[results.len() - 3])
}