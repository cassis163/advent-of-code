use std::fs;

fn main() {
    let input = fs::read_to_string("data.txt").expect("Failed to read input file");
    let calibration_values_total: u32 = get_calibration_values(&input).iter().sum();
    println!(
        "Total of all calibration values is '{}'",
        calibration_values_total.to_string()
    );
}

fn get_calibration_values(input: &String) -> Vec<u32> {
    let calibrate_line = |line: &str| {
        let digits: Vec<u32> = line.chars().filter_map(|char| char.to_digit(10)).collect();
        let digit_a = digits.first().expect("Failed to find first digit");
        let digit_b = digits.last().expect("Failed to find last digit");
        digit_a * 10 + digit_b
    };
    let filter_empty_lines = |line: &&str| !line.is_empty();

    input.lines().filter(filter_empty_lines).map(calibrate_line).collect()
}

#[cfg(test)]
mod tests {
    use crate::get_calibration_values;

    #[test]
    fn test_get_calibration_values() {
        let input = String::from(
            "1abc2\n
            pqr3stu8vwx\n
            a1b2c3d4e5f\n
            treb7uchet\n"
        );
        let actual = get_calibration_values(&input);
        let expected: Vec<u32> = vec![12, 38, 15, 77];
        assert_eq!(actual, expected);
    }
}
