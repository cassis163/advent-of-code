use std::{fs, str::Lines};

const SPELLED_OUT_NUMBERS: [(&str, u32); 9] = [
    ("one", 1),
    ("two", 2),
    ("three", 3),
    ("four", 4),
    ("five", 5),
    ("six", 6),
    ("seven", 7),
    ("eight", 8),
    ("nine", 9),
];

fn main() {
    let input: String = fs::read_to_string("data.txt").expect("Failed to read input file");
    let sanitized_lines: Vec<String> = sanitize_lines(&mut input.lines());

    let calibration_values_total: u32 = get_calibration_values_part_one(&sanitized_lines)
        .iter()
        .sum();
    println!(
        "Total of all calibration values for part one is '{}'",
        calibration_values_total.to_string()
    );

    let calibration_values_total: u32 = get_calibration_values_part_two(&sanitized_lines)
        .iter()
        .sum();
    println!(
        "Total of all calibration values for part two is '{}'",
        calibration_values_total.to_string()
    );
}

fn sanitize_lines<'a>(lines: &mut Lines) -> Vec<String> {
    lines
        .into_iter()
        .map(|s| s.replace(" ", ""))
        .filter(|s| !s.is_empty())
        .map(|s| String::from(s))
        .collect()
}

fn get_calibration_values_part_one(lines: &Vec<String>) -> Vec<u32> {
    let calibrate_line = |line: &String| {
        let digits: Vec<u32> = line.chars().filter_map(|char| char.to_digit(10)).collect();
        let digit_a = digits.first().expect("Failed to find first digit");
        let digit_b = digits.last().expect("Failed to find last digit");
        digit_a * 10 + digit_b
    };

    lines.iter().map(calibrate_line).collect()
}

fn get_calibration_values_part_two(lines: &Vec<String>) -> Vec<u32> {
    // Replace 'six' with '6' and so on
    fn transform_line(line: &String) -> String {
        let mut new_line = line.clone();
        for (word, number) in SPELLED_OUT_NUMBERS.iter() {
            let str_char = String::from(number.to_string().as_str());
            let (half_a, half_b) = word.split_at(word.len() / 2);
            let substitute = format!("{0}{1}{2}", half_a, str_char, half_b);
            new_line = new_line.replace(word, &substitute);
        }
        new_line
    }

    let transformed_lines: Vec<String> = lines.iter().map(transform_line).collect();
    // Pull the same old trick as in part one
    get_calibration_values_part_one(&transformed_lines)
}

#[cfg(test)]
mod tests {
    use crate::{get_calibration_values_part_one, get_calibration_values_part_two, sanitize_lines};

    #[test]
    fn test_get_calibration_values_part_one() {
        let input = String::from(
            "1abc2\n
            pqr3stu8vwx\n
            a1b2c3d4e5f\n
            treb7uchet",
        );
        let sanitized_lines: Vec<String> = sanitize_lines(&mut input.lines());

        let actual = get_calibration_values_part_one(&sanitized_lines);
        let expected: Vec<u32> = vec![12, 38, 15, 77];
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_get_calibration_values_part_two() {
        let input = String::from(
            "two1nine\n
            eightwothree\n
            abcone2threexyz\n
            xtwone3four\n
            4nineeightseven2\n
            zoneight234\n
            7pqrstsixteen",
        );
        let sanitized_lines: Vec<String> = sanitize_lines(&mut input.lines());

        let actual = get_calibration_values_part_two(&sanitized_lines);
        let expected: Vec<u32> = vec![29, 83, 13, 24, 42, 14, 76];
        assert_eq!(actual, expected);
    }
}
