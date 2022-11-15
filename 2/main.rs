//use std::fs;
use std::env;
use std::fs::File;
use std::io::Read;

fn main() -> std::io::Result<()> {
    let part = 3; //part 1 / 2 / 3(both)
    let samp = 1; //sample 1 / 0

    let cur_path = env::current_dir()?;
    let mut sf = "/src/sample.txt";
    if samp == 0 {
        sf = "/src/input.txt";
    }
    let file_path = cur_path.display().to_string().clone() + &sf;
    let mut input_file = File::open(file_path)?;
    let mut input_text = String::new();
    input_file.read_to_string(&mut input_text)?;

    let mut dir: Vec<&str> = Vec::new();
    let mut dist: Vec<i32> = Vec::new();
    for i in input_text.split("\n") {
        let r: Vec<&str> = i.split(" ").collect();
        //println!("{}, {}",r[0], r[1]);
        if i.len() > 2 {
            dir.push(r[0]);
            dist.push(r[1].parse().unwrap());
        }
    }
    if part == 1 {
        part_one(dir, dist);
    } else if part == 2 {
        part_two(dir, dist);
    } else {
        part_one(dir.clone(), dist.clone());
        part_two(dir.clone(), dist.clone());
    }
    Ok(())
}

fn part_one(direction: Vec<&str>, distance: Vec<i32>) {
    let mut depth = 0;
    let mut pos = 0;
    for i in 0..(direction.len()) {
        //println!("{}, {}", direction[i], distance[i]);
        if direction[i] == "forward" {
            pos += distance[i];
        } else if direction[i] == "down" {
            depth += distance[i];
        } else if direction[i] == "up" {
            depth -= distance[i];
        }
        //println!("x {}, y {}", pos, depth);
    }
    let ans = depth*pos;
    println!("Part 1: x={}, y={}, x*y={}", pos, depth, ans);
}

fn part_two(direction: Vec<&str>, distance: Vec<i32>) {
    let mut depth = 0;
    let mut pos = 0;
    let mut aim = 0;
    for i in 0..(direction.len()) {
        //println!("{}, {}", direction[i], distance[i]);
        if direction[i] == "forward" {
            pos += distance[i];
            depth += aim*distance[i];
        } else if direction[i] == "down" {
            aim += distance[i];
        } else if direction[i] == "up" {
            aim -= distance[i];
        }
        //println!("x {}, y {}", pos, depth);
    }
    let ans = depth*pos;
    println!("Part 2: x={}, y={}, x*y={}", pos, depth, ans);
}
