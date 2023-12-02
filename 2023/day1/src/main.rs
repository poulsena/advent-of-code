use std::fs;

fn part_one(input: &str) {
    let mut sum: u32 = 0;

    for line in input.lines() {
        let digits = line
            .chars()
            .filter(|c| c.is_digit(10))
            .map(|c| c.to_string())
            .collect::<Vec<String>>();

        let first_digit = digits[0].clone();
        let last_digit = digits[digits.len() - 1].clone();
        let concatenated_digits = format!("{}{}", first_digit, last_digit);

        sum += concatenated_digits.parse::<u32>().unwrap();
    }

    println!("Part one: {}", sum);
}

fn part_two(input: &str) {
    let mut sum: u32 = 0;
    let words = [
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];

    for line in input.lines() {
        let mut digits: Vec<String> = vec![];
        for (i, c) in line.chars().enumerate() {
            if c.is_digit(10) {
                digits.push(c.to_string());
            }
            for (j, word) in words.iter().enumerate() {
                if i + word.len() <= line.len() && &&line[i..i + word.len()] == word {
                    digits.push((j + 1).to_string());
                }
            }
        }
        let first_digit = digits[0].clone();
        let last_digit = digits[digits.len() - 1].clone();
        let concatenated_digits = format!("{}{}", first_digit, last_digit);

        sum += concatenated_digits.parse::<u32>().unwrap();
    }
    println!("Part one: {}", sum);
}

fn main() {
    let input = fs::read_to_string("input.txt").expect("Error reading input.txt");

    part_one(&input);
    part_two(&input);
}
