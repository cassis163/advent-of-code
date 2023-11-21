use std::fs;

fn main() {
    let report = get_report_from_file();
    let amount = get_amount_of_larger_measurements(&report);
    println!("Amount of larger measurements: {}", amount);
}

fn get_amount_of_larger_measurements(report: &[u32]) -> u32 {
    let mut amount = 0;
    for (i, x) in report.iter().enumerate().skip(1) {
        let previous_item = report.get(i - 1);
        match previous_item {
            None => continue,
            Some(previous_item) => {
                if x > previous_item {
                    amount += 1
                }
            }
        }
    }

    amount
}

fn get_report_from_file() -> Vec<u32> {
    let input = fs::read_to_string("./data/part-1.txt").unwrap();
    input
        .split("\n")
        .map(|x| x.parse::<u32>().unwrap())
        .collect()
}

#[cfg(test)]
mod tests {
    use crate::get_amount_of_larger_measurements;

    #[test]
    fn larger_measurements() {
        let report = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263];
        let result = get_amount_of_larger_measurements(&report);
        assert_eq!(result, 7);
    }
}
