use regex::Regex;
use std::fs;

struct Almanac<'a> {
    seeds: Vec<u64>,
    plans: Vec<SeedingPlan<'a>>,
}

struct SeedingPlan<'a> {
    name: &'a str,
    mappings: Vec<Mapping>,
}

struct Mapping {
    source_start: u64,
    dest_start: u64,
    range: u64,
}

impl Mapping {
    fn get_destination(&self, source: u64) -> Option<u64> {
        let is_in_range = source >= self.source_start && source <= self.source_start + self.range;
        match is_in_range {
            true => {
                let delta = source - self.source_start;
                let value = self.dest_start + delta;
                Some(value)
            }
            false => None,
        }
    }
}

impl<'a> SeedingPlan<'a> {
    fn get_destination(&self, source: u64) -> u64 {
        for mapping in &self.mappings {
            let destination = mapping.get_destination(source);
            if let Some(destination) = destination {
                return destination;
            }
        }

        source
    }
}

impl<'a> Almanac<'a> {
    fn get_seeding_plan(
        &self,
        source_name: &str,
        dest_name: &str,
    ) -> Result<&SeedingPlan, &'static str> {
        let plans = self
            .plans
            .iter()
            .filter(|plan| plan.name == format!("{}-to-{}", source_name, dest_name))
            .collect::<Vec<&SeedingPlan>>();

        if plans.len() > 1 {
            return Err("More than one seeding plan found");
        }

        let plan = *plans.first().ok_or("Failed to get plan")?;
        Ok(plan)
    }
}

fn main() {
    let input = fs::read_to_string("data.txt").unwrap();
    let sanitized_input = input.trim().to_string();
    let almanac = deserialize_input(&sanitized_input);

    let lowest_soil_number = get_lowest_soil_number_part_one(&almanac).unwrap();
    println!("Lowest soil number is {}", lowest_soil_number);

    let lowest_soil_number_part_two = get_lowest_soil_number_part_two(&almanac).unwrap();
    println!("Lowest soil number is {} (part two)", lowest_soil_number_part_two);
}

fn get_lowest_soil_number_part_one(almanac: &Almanac) -> Result<u64, &'static str> {
    let mut locations: Vec<u64> = vec![];
    for seed in almanac.seeds.iter() {
        let soil_number = get_soil_number(almanac, *seed)?;
        locations.push(soil_number);
    }

    let min = locations
        .iter()
        .min()
        .ok_or("Failed to get the lowest value")?;
    Ok(*min)
}

fn get_lowest_soil_number_part_two(almanac: &Almanac) -> Result<u64, &'static str> {
    let seed_pairs = get_seed_pairs(almanac);

    let mut locations: Vec<u64> = vec![];
    for (start_seed, range) in seed_pairs.iter() {
        let seeds = **start_seed..(*start_seed + *range);
        for seed in seeds {
            let soil_number = get_soil_number(almanac, seed)?;
            locations.push(soil_number);
        }
    }

    let min = locations
        .iter()
        .min()
        .ok_or("Failed to get the lowest value")?;
    Ok(*min)
}

fn get_seed_pairs<'a>(almanac: &'a Almanac) -> Vec<(&'a u64, &'a u64)> {
    almanac
        .seeds
        .iter()
        .enumerate()
        .filter_map(|(i, seed)| {
            if (i + 1) % 2 == 0 {
                return None;
            }

            let next_seed = almanac
                .seeds
                .get(i + 1)
                .expect("Failed to get second seed in pair");
            Some((seed, next_seed))
        })
        .collect()
}

fn get_soil_number(almanac: &Almanac, seed: u64) -> Result<u64, &'static str> {
    let soil = almanac
        .get_seeding_plan("seed", "soil")?
        .get_destination(seed);
    let fertilizer = almanac
        .get_seeding_plan("soil", "fertilizer")?
        .get_destination(soil);
    let water = almanac
        .get_seeding_plan("fertilizer", "water")?
        .get_destination(fertilizer);
    let light = almanac
        .get_seeding_plan("water", "light")?
        .get_destination(water);
    let temperature = almanac
        .get_seeding_plan("light", "temperature")?
        .get_destination(light);
    let humidity = almanac
        .get_seeding_plan("temperature", "humidity")?
        .get_destination(temperature);
    let location = almanac
        .get_seeding_plan("humidity", "location")?
        .get_destination(humidity);
    Ok(location)
}

fn deserialize_input(sanitized_input: &String) -> Almanac {
    let groups = sanitized_input.split("\n\n").collect::<Vec<&str>>();

    let seeds_str = groups.get(0).expect("Failed to obtain seeds");
    let seeds = get_numbers(&seeds_str.to_string());

    let map_groups = groups[1..].to_vec();
    let plans: Vec<SeedingPlan> = map_groups
        .iter()
        .map(|map_group| {
            let lines = map_group.lines().collect::<Vec<&str>>();
            let name_line = lines.get(0).unwrap();
            let name = get_name(name_line).expect("Failed to get name");

            let number_lines = lines[1..].iter().collect::<Vec<&&str>>();
            let mappings: Vec<Mapping> = number_lines
                .iter()
                .map(|line| {
                    let numbers = get_numbers(&line.to_string());
                    let dest_start = numbers.get(0).expect("Failed to get destination start");
                    let source_start = numbers.get(1).expect("Failed to get source start");
                    let range = numbers.get(2).expect("Failed to get range");

                    Mapping {
                        dest_start: *dest_start,
                        range: *range,
                        source_start: *source_start,
                    }
                })
                .collect();

            SeedingPlan { name, mappings }
        })
        .collect();

    Almanac { seeds, plans }
}

fn get_name(s: &str) -> Option<&str> {
    let re = Regex::new(r"\b(\w+-to-\w+)\b").unwrap();
    re.captures_iter(s)
        .map(|c| c.extract::<1>().0)
        .collect::<Vec<&str>>()
        .first()
        .map(|v| *v)
}

fn get_numbers(s: &String) -> Vec<u64> {
    let re = Regex::new(r"(\d+)").unwrap();
    re.captures_iter(s)
        .map(|c| {
            let value = c.extract::<1>().0;
            value.parse::<u64>().unwrap()
        })
        .collect()
}

#[cfg(test)]
mod tests {
    use crate::{
        deserialize_input, get_lowest_soil_number_part_one, get_lowest_soil_number_part_two,
    };

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
    fn test_get_lowest_soil_number_part_one() {
        let sanitized_input = INPUT.trim().to_string();
        let almanac = deserialize_input(&sanitized_input);

        let actual = get_lowest_soil_number_part_one(&almanac).unwrap();
        let expected: u64 = 35;
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_get_lowest_soil_number_part_two() {
        let sanitized_input = INPUT.trim().to_string();
        let almanac = deserialize_input(&sanitized_input);

        let actual = get_lowest_soil_number_part_two(&almanac).unwrap();
        let expected: u64 = 46;
        assert_eq!(actual, expected);
    }
}
