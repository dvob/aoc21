use std::{str::FromStr, fmt::Display};

type StdResult<T> = Result<T, Box<dyn std::error::Error>>;

const NEW_FISH_CYCLE: u32 = 6;

struct Fish {
    first_cycle: bool,
    timer: u32,
}

impl Fish {
    fn new() -> Self {
        Self {
            timer: NEW_FISH_CYCLE + 2,
            first_cycle: true,
        }
    }
    fn new_with_timer(timer: u32) -> Self {
        Self {
            timer: timer,
            first_cycle: false,
        }
    }

    fn next_day(&mut self) -> bool {
        if self.timer == 0 {
            self.timer = NEW_FISH_CYCLE;
            self.first_cycle = false;
            true
        } else {
            self.timer -= 1;
            false
        }
    }
}

// 1 -> add fishes here
// 0 -> move fishes here

// 6 ->
// 5 ->
// 4 -> 
// 3 ->
// 2 ->
// 1 ->
// 0 ->
#[derive(Debug)]
struct FishPool2 {
    buckets: [u64; 9],
}
impl FishPool2 {
    fn new() -> Self {
        Self {
            buckets: [0; 9],
        }
    }

    fn next_day(&mut self) {
        self.buckets.rotate_left(1);
        self.buckets[6] += self.buckets[8];
    }

    fn num_fishes(&self) -> u64 {
        self.buckets.iter().sum()
    }
}

impl Display for FishPool2 {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.num_fishes())
    }
}

impl FromStr for FishPool2 {
    type Err = Box<dyn std::error::Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut fp = FishPool2::new();
        for timer in s.split(",") {
            let num: usize = timer.parse()?;
            fp.buckets[num] += 1
        }
        Ok(fp)
    }
}

struct FishPool(Vec<Fish>);

impl FishPool {
    fn next_day(&mut self) {
        let mut new_fishes = 0;
        for fish in self.0.iter_mut() {
            if fish.next_day() {
                new_fishes += 1
            }
        }
        for i in 0..new_fishes {
            self.0.push(Fish::new())
        }
    }
    fn num_fishes(&self) -> usize {
        self.0.len()
    }
}

impl Display for FishPool {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut iter = self.0.iter();
        let first = iter.next();
        match first {
            Some(fish) => write!(f, "{}", fish.timer)?,
            None => write!(f, "no fishes")?,
        };
        for fish in iter {
            write!(f, ",{}", fish.timer)?;
        }
        Ok(())
    }
}

impl FromStr for FishPool {
    type Err = Box<dyn std::error::Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut fishes = vec![];
        for timer in s.split(",") {
            fishes.push(Fish::new_with_timer(timer.parse()?));
        }
        Ok(FishPool(fishes))
    }
}

fn main() -> StdResult<()> {
    let input = std::fs::read_to_string("input")?;
    let mut fp: FishPool2 = input.trim().parse()?;
    for _i in 1..=256 {
        fp.next_day()
    }
    println!("{}", fp.num_fishes());
    Ok(())
}

#[cfg(test)]
mod tests {
    use crate::{StdResult, FishPool, FishPool2};

    #[test]
    fn test() -> StdResult<()> {
        let input = "3,4,3,1,2";

        let mut fp: FishPool2 = input.parse()?;
        println!("Initial state: {}", fp);
        for i in 1..=18 {
            fp.next_day();
            println!("After {} days: {:?} {}", i, fp, fp);
        }
        assert_eq!(fp.num_fishes(), 26);
        for i in 1..=(80-18) {
            fp.next_day()
        }
        assert_eq!(fp.num_fishes(), 5934);
        for i in 1..=(256-80) {
            fp.next_day()
        }
        assert_eq!(fp.num_fishes(), 26984457539);
        Ok(())
    }
}