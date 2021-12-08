mod input;

// sort data first please!!!
fn minimal_cost(data: &Vec<i64>, cost_fn: fn(&Vec<i64>, i64) -> i64) -> i64 {
	let left = data[0];
	let right = data[data.len() - 1];
	let mut cost = cost_fn(data, left);
	for p in (left + 1)..right {
		let c = cost_fn(data, p);
		if c > cost {
			break;
		};
		cost = c;
	}
	cost
}

fn cost1(data: &Vec<i64>, pos: i64) -> i64 {
	data.iter().map(|x| (x - pos).abs()).sum()
}

fn cost2(data: &Vec<i64>, pos: i64) -> i64 {
	data.iter()
		.map(|x| {
			let d = (x - pos).abs();
			d * (d + 1) / 2
		})
		.sum()
}

fn part1(data: &Vec<i64>) -> i64 {
	minimal_cost(data, cost1)
}

fn part2(data: &Vec<i64>) -> i64 {
	minimal_cost(data, cost2)
}

fn main() {
	// let mut sample: Vec<i64> = include_str!("sample")
	// 	.split(',')
	// 	.map(|n| n.parse().unwrap())
	// 	.collect();
	let mut sample = Vec::from(input::SAMPLE);
	sample.sort();
	let mut input = Vec::from(input::INPUT);
	input.sort();

	println!("=== Part 1 ===");
	println!("sample : {}", part1(&sample));
	println!("input  : {}", part1(&input));
	println!("=== Part 2 ===");
	println!("sample : {}", part2(&sample));
	println!("input  : {}", part2(&input));

	// === Part 1 ===
	// sample : 37
	// input  : 340987
	// === Part 2 ===
	// sample : 168
	// input  : 96987874
}
