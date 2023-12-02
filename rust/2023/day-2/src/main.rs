use std::fs;

struct Game {
    id: u32,
    bags: Vec<Bag>,
}

struct Bag {
    red_amount: u32,
    green_amount: u32,
    blue_amount: u32,
}

fn main() {
    let input = fs::read_to_string("data.txt").expect("Failed to read file");
    let bag_config = Bag {
        red_amount: 12,
        green_amount: 13,
        blue_amount: 14,
    };

    let possible_games_count = count_possible_games(&input, &bag_config);
    println!(
        "{} amount of games possible with this configuration",
        possible_games_count
    );

    let power_sets_sum = get_power_sets_sum(&input);
    println!("Sum of power sets is {}", power_sets_sum);
}

fn get_power_sets_sum(input: &String) -> u32 {
    let game_strs: Vec<&str> = input.split("\n").filter(|line| !line.is_empty()).collect();
    let games: Vec<Game> = game_strs
        .iter()
        .map(|s| {
            let game_str = String::from(*s);
            get_game_from_line(&game_str)
        })
        .collect();
    let min_bags: Vec<Bag> = games.iter().map(|game| {
        Bag {
            red_amount: game.bags.iter().map(|bag| bag.red_amount).max().unwrap(),
            blue_amount: game.bags.iter().map(|bag| bag.blue_amount).max().unwrap(),
            green_amount: game.bags.iter().map(|bag| bag.green_amount).max().unwrap(),
        }
    }).collect();

    min_bags.iter().map(|bag| {
        bag.red_amount * bag.green_amount * bag.blue_amount
    }).sum()
}

fn count_possible_games(input: &String, bag_config: &Bag) -> u32 {
    let game_strs: Vec<&str> = input.split("\n").filter(|line| !line.is_empty()).collect();
    let games: Vec<Game> = game_strs
        .iter()
        .map(|s| {
            let game_str = String::from(*s);
            get_game_from_line(&game_str)
        })
        .collect();

    let possible_games: Vec<&Game> = games
        .iter()
        .filter(|game| {
            !game.bags.iter().any(|bag| {
                bag.red_amount > bag_config.red_amount
                    || bag.green_amount > bag_config.green_amount
                    || bag.blue_amount > bag_config.blue_amount
            })
        })
        .collect();
    possible_games.iter().map(|game| game.id).sum()
}

fn get_game_from_line(line: &String) -> Game {
    let sanitized_line = line.trim().replace(" ", "");
    let game_and_bags_str: Vec<&str> = sanitized_line.split(":").collect();

    let game_str = game_and_bags_str.get(0).expect("Game is not present");
    let game_id = extract_number(game_str);

    let bags_str = game_and_bags_str.get(1).expect("Bags are not present");
    let bag_strs: Vec<&str> = bags_str.split(";").collect();

    let bags: Vec<Bag> = bag_strs.iter().map(create_bag_from_str).collect();

    Game { id: game_id, bags }
}

fn create_bag_from_str(bag_str: &&str) -> Bag {
    let cube_strs: Vec<&str> = bag_str.split(",").collect();
    let mut bag = Bag {
        red_amount: 0,
        green_amount: 0,
        blue_amount: 0,
    };

    for cube_str in cube_strs {
        if cube_str.ends_with("red") {
            bag.red_amount += extract_number(cube_str);
        }

        if cube_str.ends_with("green") {
            bag.green_amount += extract_number(cube_str);
        }

        if cube_str.ends_with("blue") {
            bag.blue_amount += extract_number(cube_str);
        }
    }

    bag
}

fn extract_number(s: &str) -> u32 {
    let result: String = s
        .chars()
        .skip_while(|c| !c.is_numeric())
        .take_while(|c| c.is_numeric())
        .collect();
    result.parse().unwrap()
}

#[cfg(test)]
mod tests {
    use crate::{count_possible_games, Bag, get_power_sets_sum};

    #[test]
    fn test_count_possible_games() {
        let input = String::from(
            "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n
            Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n
            Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n
            Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n
            Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
        );

        let bag_config = Bag {
            red_amount: 12,
            green_amount: 13,
            blue_amount: 14,
        };
        let actual = count_possible_games(&input, &bag_config);
        let expected = 8;
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_get_power_sets_sum() {
        let input = String::from(
            "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n
            Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n
            Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n
            Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n
            Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
        );

        let actual = get_power_sets_sum(&input);
        let expected = 2286;
        assert_eq!(actual, expected);
    }
}
