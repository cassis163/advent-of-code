use std::fs;

use regex::Regex;

struct Card {
    // id: u32,
    winning_numbers: Vec<u32>,
    owned_numbers: Vec<u32>,
}

impl Card {
    fn get_value(&self) -> u32 {
        self.winning_numbers.iter().fold(0, |acc, n| {
            if !self.owned_numbers.contains(n) {
                return acc;
            }

            match acc {
                0 => 1,
                _ => acc * 2,
            }
        })
    }
}

fn main() {
    let input = fs::read_to_string("data.txt").expect("Failed to read file");
    let sanitized_input = sanitize_input(&input);

    let total_points: u32 = get_total_points(&sanitized_input);
    println!("Total points of cards is {}", total_points);
}

fn sanitize_input(input: &String) -> String {
    input.trim().to_string()
}

fn extract_numbers(input: &str) -> Vec<u32> {
    let re = Regex::new(r"(\d+)").unwrap();
    re.captures_iter(input).map(|c| c.extract::<1>()).map(|c| {
        c.0.parse::<u32>().unwrap()
    }).collect()
}

fn get_total_points(input: &String) -> u32 {
    let cards = get_cards_from_input(&input);
    cards.iter().map(Card::get_value).sum()
}

fn get_cards_from_input(input: &String) -> Vec<Card> {
    input
        .lines()
        .filter_map(|line| {
            if line.trim().is_empty()  {
                return None;
            }

            let card_and_number_halves: Vec<&str> = line.split(":").collect();

            // let card_half = card_and_number_halves
            //     .get(0)
            //     .expect("Card half is not present");
            // let card_id = *extract_numbers(&card_half)
            //     .first()
            //     .expect("Failed to find card ID");

            let numbers_half = card_and_number_halves
                .get(1)
                .expect("Numbers half is not present");
            let winning_and_owned_numbers: Vec<&str> = numbers_half.split("|").collect();
            let winning_numbers = winning_and_owned_numbers
                .get(0)
                .expect("Failed to get winning numbers");
            let owned_numbers = winning_and_owned_numbers
                .get(1)
                .expect("Failed to get owned numbers");

            let winning_numbers = extract_numbers(&winning_numbers);
            let owned_numbers = extract_numbers(&owned_numbers);

            Some(Card {
                // id: card_id,
                winning_numbers,
                owned_numbers,
            })
        })
        .collect()
}

#[cfg(test)]
mod tests {
    use crate::get_total_points;

    #[test]
    fn test_get_total_points() {
        let input = String::from(
            "
            Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n
            Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n
            Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n
            Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n
            Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n
            Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
            ",
        );

        let actual = get_total_points(&input);
        let expected = 13;
        assert_eq!(actual, expected);
    }
}
