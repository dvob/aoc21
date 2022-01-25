use std::str::Chars;

type StdResult<T> = Result<T, Box<dyn std::error::Error>>;

fn main() -> StdResult<()> {
    let input = std::fs::read_to_string("input")?;
    let mut err_score: u32 = 0;
    let mut scores = vec![];
    for line in input.lines() {
        match check_line(&mut line.chars()) {
            Err(c) => err_score += get_points(c),
            Ok(score) => if score != 0 {scores.push(score)},
        }
    }
    println!("err score: {}", err_score);
    scores.sort();
    let index = scores.len() / 2;
    println!("score {}", scores[index]);
    Ok(())
}

fn check_line(line: &mut Chars) -> Result<u64, char> {
    let open = match line.next() {
        None => return Ok(0),
        Some(c) => c,
    };
    if !is_open(open) {
        return Err(open)
    }
    let mut score: u64 = 0;
    loop {
        let mut peek = line.clone();
        match peek.next() {
            None => break,
            Some(c) => {
                if is_open(c) {
                    score = check_line(line)?;
                } else {
                    break
                }
            }
        };
    }
    score *= 5;
    match line.next() {
        None => Ok(score + get_points2(open)),
        Some(close) => {
            if !match_paren(open, close) {
                return Err(close)
            } else {
                Ok(0)
            }
        }
    }
}

fn match_paren(open: char, close: char) -> bool {
    match (open, close) {
        ('(', ')') => true,
        ('[', ']') => true,
        ('{', '}') => true,
        ('<', '>') => true,
        _ => false,
    }
}

fn is_open(c: char) -> bool {
    match c {
        '(' => true,
        '[' => true,
        '{' => true,
        '<' => true,
        _ => false,
    }
}

fn get_points2(c: char) -> u64 {
    match c {
        '(' => 1,
        '[' => 2,
        '{' => 3,
        '<' => 4,
        _ => panic!("invalid character {}", c),
    }
}
fn get_points(c: char) -> u32 {
    match c {
        ')' => 3,
        ']' => 57,
        '}' => 1197,
        '>' => 25137,
        _ => panic!("invalid character {}", c),
    }
}
