use std::fs::File;
use std::io::{self, BufRead};

fn main() {
    use std::time::Instant;

    let mut start = Instant::now();
    let mut result;
    {
        result = part1();
    }
    
    let mut elapsed = start.elapsed();
    match result {
        Ok(s) => println!("Part 1 {} in {:.4?}\n", s, elapsed),
        Err(err) => panic!("Failure {}\n", err)
    }

    start = Instant::now();
    result = part2();
    elapsed = start.elapsed();
    match result {
        Ok(s) => println!("part 2 {} in {:.4?}\n", s, elapsed),
        Err(err) => panic!("Failure {}\n", err)
    }
}

fn part1() -> Result<u32, Box<dyn std::error::Error>> {
    let file = File::open("./input.txt")?;
    let reader = io::BufReader::new(file);
    let mut result: u32 = 0;

    for line in reader.lines() {

        match line {
            Ok(s) => {
                let len = s.len() / 2;
                let mut first_bag = s.into_bytes();
                let mut second_bag = first_bag.split_off(len);
                second_bag.sort_unstable();
                first_bag.sort_unstable();
                let m = find_match(second_bag, first_bag);
                result += get_priority(m) as u32;
            }
            Err(err) => panic!("{}", err)
        }
    }
    Ok(result)
}

fn part2() -> Result<u32, Box<dyn std::error::Error>> {
    let file = File::open("./input.txt")?;
    let reader = io::BufReader::new(file);
    let mut result: u32 = 0;
    let mut lines = reader.lines();

    while let (Some(l1), Some(l2), Some(l3)) = (lines.next(), lines.next(), lines.next()) {
        let mut bag1 = l1.unwrap().into_bytes();
        let mut bag2 = l2.unwrap().into_bytes();
        let mut bag3 = l3.unwrap().into_bytes();

        bag1.sort_unstable();
        bag2.sort_unstable();
        bag3.sort_unstable();

        let dupes = get_dupes(bag1, bag2);
        let all_dupes = get_dupes(dupes, bag3);
        
        if all_dupes.len() != 1 {
            panic!("There were {} matches in the final dupes", all_dupes.len());
        }

        result += get_priority(all_dupes[0]) as u32;
    }
    
    Ok(result)
}

fn get_dupes(first: Vec<u8>, second: Vec<u8>) -> Vec<u8> {
    let mut dupes: Vec<u8> = vec![];

    let mut first_idx = 0;
    let mut second_idx = 0;

    while first_idx < first.len() && second_idx < second.len(){
        if first[first_idx] == second[second_idx] {
            dupes.push(first[first_idx]);
            first_idx += 1;
            while first_idx < first.len() && first[first_idx] == first[first_idx - 1]  {
                // Skip adding the same value multiple times
                first_idx += 1;
            }
        } else if first[first_idx] < second[second_idx] {
            first_idx += 1;
        } else {
            second_idx += 1;
        }
    }

    dupes
}

fn find_match(first: Vec<u8>, second: Vec<u8>) -> u8 {
    /*  Comparing all values is going to be O(N^2), however if we first sort the values,
        then we can move an index through each vector which is O(N) and the sorting should
        be O(N log(N)).
    */
    let mut first_idx = 0;
    let mut second_idx = 0;

    while first_idx < first.len() && second_idx < second.len() {
        if first[first_idx] == second[second_idx] {
            return first[first_idx];
        } else if first[first_idx] < second[second_idx] {
            first_idx += 1;
        } else {
            second_idx += 1;
            if second_idx >= second.len() {
                panic!("no match found\n");
            }
        }
    }
    panic!("no match found");
}

fn get_priority(item: u8) -> u8 {
    
    let lower_case_base = 'a' as u8 - 1;
    let upper_case_base = 'A' as u8 - 1;
    let result;
    if item > lower_case_base {
        result = item - lower_case_base
    } else {
        result = item - upper_case_base + 26
    }
    result
}
