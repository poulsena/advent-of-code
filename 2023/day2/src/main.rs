use std::fs;

fn part_one(input: &str) {
    const MAX_RED: usize = 12;
    const MAX_GREEN: usize = 13;
    const MAX_BLUE: usize = 14;

    let mut total: usize = 0;

    for (i, line) in input.lines().enumerate() {
        let mut count = true;
        let id: usize = i + 1;
        let sets = line.split(": ").collect::<Vec<&str>>()[1]
            .split("; ")
            .collect::<Vec<&str>>();

        'set: for set in sets {
            let mut sum_red: usize = 0;
            let mut sum_green: usize = 0;
            let mut sum_blue: usize = 0;
            let set = set.split(", ").collect::<Vec<&str>>();

            for color in set {
                let digits: usize = color
                    .chars()
                    .filter(|c| c.is_digit(10))
                    .collect::<String>()
                    .parse()
                    .unwrap();

                let color_string = color
                    .trim()
                    .chars()
                    .filter(|c| c.is_alphabetic())
                    .collect::<String>();

                match color_string.as_str() {
                    "red" => {
                        if sum_red + digits <= MAX_RED {
                            sum_red += digits;
                        } else {
                            count = false;
                            break 'set;
                        }
                    }
                    "green" => {
                        if sum_green + digits <= MAX_GREEN {
                            sum_green += digits;
                        } else {
                            count = false;
                            break 'set;
                        }
                    }
                    "blue" => {
                        if sum_blue + digits <= MAX_BLUE {
                            sum_blue += digits;
                        } else {
                            count = false;
                            break 'set;
                        }
                    }
                    _ => {
                        count = false;
                        break 'set;
                    }
                }
            }
        }

        if count {
            total += id;
        }
    }

    println!("Part One: Total: {}", total);
}

fn main() {
    let input = fs::read_to_string("input.txt").expect("Error reading input.txt");

    part_one(&input);
}
