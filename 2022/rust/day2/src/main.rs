use std::fs::File;
use std::io::{self, BufRead};

#[derive(PartialEq, Eq, Copy, Clone)]
enum HandShape {
    Rock = 1,
    Paper,
    Scissors
}

impl HandShape {
    fn from_i8(v: i8) -> HandShape {
        match v {
            0 => HandShape::Scissors,
            1 => HandShape::Rock,
            2 => HandShape::Paper,
            3 => HandShape::Scissors,
            4 => HandShape::Rock,
            _ => panic!("error i8 to hand shape {}", v)
        }

    }

    fn from_char(c: char) -> HandShape {
        match c {
            'A' => HandShape::Rock,
            'X' => HandShape::Rock,
            'B' => HandShape::Paper,
            'Y' => HandShape::Paper,
            'C' => HandShape::Scissors,
            'Z' => HandShape::Scissors,
            _ => panic!("error")
        }
    }
}



#[derive(PartialEq, Eq)]
enum HandResult {
    Win,
    Lose,
    Draw
}

impl HandResult {
    fn from_char(c: char) -> HandResult {
        match c {
            'X' => HandResult::Lose,
            'Y' => HandResult::Draw,
            'Z' => HandResult::Win,
            _ => panic!("error"),
        }
    }
}

fn get_hand_result(other: HandShape, my: HandShape) -> HandResult {
    // Did we tie (this covers all possible ties)
    if other == my {
        return HandResult::Draw;
    }

    // Did we win
    // This works for other == Rock or Paper
    if my as u8 == (other as u8 + 1) || ((other == HandShape::Scissors) && (my == HandShape::Rock)) {
        HandResult::Win
    } else {
        HandResult::Lose
    }    
} 

fn get_shape_for_result(other: HandShape, result: HandResult) -> HandShape {
    if result == HandResult::Draw {
        return other;
    }

    let mut v = other as i8;
    if result == HandResult::Win {
        v = v + 1; 
    } else if result == HandResult::Lose {
        v = v - 1;
    }
    return HandShape::from_i8(v);
}

fn score_hand(other: HandShape, my: HandShape) -> u32 {
    let score = my as u32;
    let result = get_hand_result(other, my);

    if result == HandResult::Lose {
        score
    } else if result == HandResult::Draw {
        score + 3
    } else {
        score + 6
    }
}

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
    let file = File::open("./input.txt")?;
    let reader = io::BufReader::new(file);
    let mut result: u32 = 0;

    for line in reader.lines() {

        match line {
            Ok(s) => {
                let p1 = HandShape::from_char(s.chars().nth(0).unwrap());
                let p2 = HandShape::from_char(s.chars().nth(2).unwrap());
                let current = score_hand(p1, p2);
                println!("{}", current);
                result += current;
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

    for line in reader.lines() {

        match line {
            Ok(s) => {
                let p1 = HandShape::from_char(s.chars().nth(0).unwrap());
                let shape = HandResult::from_char(s.chars().nth(2).unwrap());
                let my = get_shape_for_result(p1, shape);
                let current = score_hand(p1, my);
                println!("{}", current);
                result += current;
            }
            Err(err) => panic!("{}", err)
        }
    }
    Ok(result)
}
