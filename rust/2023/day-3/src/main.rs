use std::fs;

fn main() {
    let input = fs::read_to_string("data.txt").expect("Failed to read file");
    let sanitized_input = sanitize_input(&input);

    // This solution sucks. You can probably do something clever and concise here with string slices.
    let sum: u32 = get_numbers_adjacent_to_symbol(&sanitized_input)
        .iter()
        .sum();
    println!("Sum of numbers adjacent to a symbol is {}", sum);
}

fn get_numbers_adjacent_to_symbol(input: &String) -> Vec<u32> {
    let board = deserialise_input(input);
    let mut numbers: Vec<u32> = vec![];

    let mut current_number_str = String::from("");
    let mut has_symbol_neighbour = false;
    for (y, row) in board.iter().enumerate() {
        for (x, cell) in row.iter().enumerate() {
            if let Some(char) = cell {
                let next_cell_exists = row.get(x + 1).is_some();
                let next_cell_is_digit = row
                    .get(x + 1)
                    .map_or(false, |cell| cell.map_or(false, |char| char.is_digit(10)));
                if char.is_digit(10) {
                    current_number_str.push(*char);
                    if !has_symbol_neighbour {
                        has_symbol_neighbour = has_neighbour_symbol(&board, x, y);
                    }
                }

                if !next_cell_is_digit && has_symbol_neighbour {
                    let number = current_number_str.parse::<u32>().unwrap();
                    numbers.push(number);
                    current_number_str = String::from("");
                    has_symbol_neighbour = false;
                }

                if !next_cell_exists {
                    current_number_str = String::from("");
                    has_symbol_neighbour = false;
                }
            } else {
                current_number_str = String::from("");
                has_symbol_neighbour = false;
            }
        }
    }

    numbers
}

fn check_square<'a>(
    board: &'a Vec<Vec<Option<char>>>,
    x: usize,
    y: usize,
) -> Option<&'a Option<char>> {
    let row = board.get(y)?;
    row.get(x)
}

fn is_square_symbol<'a>(board: &'a Vec<Vec<Option<char>>>, x: usize, y: usize) -> bool {
    check_square(board, x, y).map_or(false, |row| row.map_or(false, |cell| !cell.is_digit(10)))
}

fn has_neighbour_symbol<'a>(board: &'a Vec<Vec<Option<char>>>, x: usize, y: usize) -> bool {
    let x_start = match x {
        0 => 0,
        _ => x - 1,
    };
    for x in x_start..(x + 2) {
        let y_start = match y {
            0 => 0,
            _ => y - 1,
        };
        for y in y_start..(y + 2) {
            if is_square_symbol(board, x, y) {
                return true;
            }
        }
    }

    false
}

fn deserialise_input(input: &String) -> Vec<Vec<Option<char>>> {
    let map_chars = |c: char| -> Option<char> {
        if c == '.' {
            return None;
        }

        Some(c)
    };

    input
        .lines()
        .filter(|line| !line.is_empty())
        .map(|line| line.chars().map(map_chars).collect::<Vec<Option<char>>>())
        .collect()
}

fn sanitize_input(input: &String) -> String {
    input.trim().replace(" ", "")
}

#[cfg(test)]
mod tests {
    use crate::{get_numbers_adjacent_to_symbol, sanitize_input};

    #[test]
    fn test_get_numbers_adjacent_to_symbol() {
        let input = String::from(
            "467..114..\n
            ...*......\n
            ..35..633.\n
            ......#...\n
            617*......\n
            .....+.58.\n
            ..592.....\n
            ......755.\n
            ...$.*....\n
            .664.598..",
        );
        let sanitized_input = sanitize_input(&input);

        let actual = get_numbers_adjacent_to_symbol(&sanitized_input);
        let expected = vec![467, 35, 633, 617, 592, 755, 664, 598];
        expected.iter().for_each(|v| {
            let does_contain = actual.contains(&v);
            assert!(does_contain);
        });
    }
}
