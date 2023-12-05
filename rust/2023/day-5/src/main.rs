use regex::Regex;
use std::{collections::HashMap, fs};

struct Almanac<'a> {
    seeds: Vec<u32>,
    maps: Vec<SourceToDestMap<'a>>,
}

struct SourceToDestMap<'a> {
    source_name: &'a str,
    dest_name: &'a str,
    source_to_dest: HashMap<&'a u32, &'a u32>,
}

impl<'a> SourceToDestMap<'a> {
    fn get_full_name(&self) -> String {
        format!("{}-to-{}", self.source_name, self.dest_name)
    }
}

fn main() {
    let input = fs::read_to_string("data.txt").unwrap();
    deserialize_input(&input);
}

fn deserialize_input(input: &String) -> Result<Almanac, &'static str> {
    let sanitized_input = input.trim().to_string();
    let groups = sanitized_input.split("\n\n").collect::<Vec<&str>>();

    let seeds_str = groups.get(0).ok_or("Failed to obtain seeds")?;
    let seeds = get_numbers(&seeds_str.to_string());

    let map_groups = groups[1..].to_vec();
    let maps: Vec<SourceToDestMap> = map_groups.iter().map(|map_group| {
        let lines = map_group.lines().collect::<Vec<&str>>();
        let name_line = lines.get(0).unwrap();
        let source_name = get_source_name(name_line);
        let dest_name = get_dest_name(name_line);

        let number_lines = lines[1..].iter().collect::<Vec<>>();
    }).collect();

    Ok(Almanac {
        seeds,
        maps
    })
}

fn get_source_name(name_line: &str) -> &str {
    todo!()
}

fn get_dest_name(name_line: &str) -> &str {
    todo!()
}

fn get_numbers(s: &String) -> Vec<u32> {
    let number_re = Regex::new(r"(\d+)").unwrap();
    number_re
        .captures_iter(s)
        .map(|c| {
            let value = c.extract::<1>().0;
            value.parse::<u32>().unwrap()
        })
        .collect()
}

#[cfg(test)]
mod tests {
    const INPUT: &str = "
    seeds: 79 14 55 13

    seed-to-soil map:
    50 98 2
    52 50 48

    soil-to-fertilizer map:
    0 15 37
    37 52 2
    39 0 15

    fertilizer-to-water map:
    49 53 8
    0 11 42
    42 0 7
    57 7 4

    water-to-light map:
    88 18 7
    18 25 70

    light-to-temperature map:
    45 77 23
    81 45 19
    68 64 13

    temperature-to-humidity map:
    0 69 1
    1 0 69

    humidity-to-location map:
    60 56 37
    56 93 4
    ";

    #[test]
    fn test() {}
}
