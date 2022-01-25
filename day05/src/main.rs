use std::{fmt::Display, str::FromStr, ops::Range};

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut f = Field::new(1000);
    let input = std::fs::read_to_string("input")?;
    let lines = read_entries(input.as_str())?;
    f.draw_lines(lines);
    println!("result: {}", f.result());
    //println!("{}", f);

    Ok(())
}

fn read_entries(input: &str) -> Result< Vec<Line<usize>>, Box<dyn std::error::Error>> {
    let mut entries = vec![];
    for line in input.lines() {
        entries.push(line.parse()?);
    };
    Ok(entries)
}

#[derive(Debug)]
struct Line<T: FromStr + std::fmt::Debug> {
    from: Point<T>,
    to: Point<T>,
}

impl<T: FromStr + std::fmt::Debug> Line<T> {
    fn new(x1: T, y1: T, x2: T, y2: T) -> Line<T> {
        Self {
            from: Point { x: x1, y: y1 },
            to: Point { x: x2, y: y2 },
        }
    }
}

impl<T: FromStr + std::fmt::Debug> FromStr for Line<T> {
    type Err = Box<dyn std::error::Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (from, to) = s.split_once(" -> ").ok_or("missing delimiter")?;
        Ok(Line{
            from: from.parse()?,
            to: to.parse()?,
        })
    }
}

#[derive(Debug,Clone,Copy)]
struct Point<T: std::fmt::Debug> {
    x: T,
    y: T,
}

impl<T: std::fmt::Debug> FromStr for Point<T> where T: FromStr {
    type Err = Box<dyn std::error::Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (x, y) = s.split_once(",").ok_or("missing delimiter")?;
        Ok(Point{
            x: x.parse().map_err(|_| "failed to parse x")?,
            y: y.parse().map_err(|_| "failed to parse y")?,
        })
    }
}

#[derive(Debug)]
struct Field(Vec<Vec<usize>>);

impl Field {
    fn new(size: usize) -> Self {
        Field(vec![vec![0; size]; size])
    }

    fn draw_lines(&mut self, lines: Vec<Line<usize>>) {
        for line in lines {
            if !(line.from.x == line.to.x || line.from.y == line.to.y) {
                self.draw_diagonal(line.from, line.to);
            } else if  line.from.x == line.to.x {
                self.draw_y(line.from.x, Field::get_range(line.from.y, line.to.y))
            } else {
                self.draw_x(line.from.y, Field::get_range(line.from.x, line.to.x))
            }
        }
    }

    fn result(&self) -> usize {
        self.0.iter().flat_map(|i| i.iter()).filter(|i | **i >= 2).count()
    }

    fn draw_diagonal(&mut self, a: Point<usize>, b: Point<usize>) {
        // OP: do not allocate here
        let x = if a.x < b.x {
            (a.x..=b.x).collect::<Vec<usize>>()
        } else {
            (b.x..=a.x).rev().collect::<Vec<usize>>()
        };
        let y = if a.y < b.y {
            (a.y..=b.y).collect::<Vec<usize>>()
        } else {
            (b.y..=a.y).rev().collect::<Vec<usize>>()
        };
        for i in 0..x.len() {
            self.0[y[i]][x[i]] += 1
        }
    }

    fn get_range(a: usize, b: usize) -> Range<usize> {
        if a > b {
            b..a+1
        } else {
            a..b+1
        }
    }

    fn draw_x(&mut self, x: usize, range: Range<usize>) {
        for i in range {
            self.0[x][i] += 1;
        }
    }
    fn draw_y(&mut self, y: usize, range: Range<usize>) {
        for i in range {
            self.0[i][y] += 1;
        }
    }
}

impl Display for Field {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for row in &self.0 {
            for col in row {
                if *col == 0 {
                    write!(f, ".")?;
                } else {
                    write!(f, "{}", col)?;
                }
            }
            write!(f, "\n")?;
        }
        return Ok(())
    }
}

#[cfg(test)]
mod tests {
    use crate::{read_entries, Field};

    #[test]
    fn test_part1() -> Result<(), Box<dyn std::error::Error>> {
        let expected_output = r#"1.1....11.
.111...2..
..2.1.111.
...1.2.2..
.112313211
...1.2....
..1...1...
.1.....1..
1.......1.
222111....
"#;
        let input = r#"0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2"#;
        let lines = read_entries(input)?;
        let mut f = Field::new(10);
        f.draw_lines(lines);
        let output = format!("{}", f);
        if output != expected_output {
            println!("-- expected --");
            println!("'{}'", expected_output);
            println!("-- actual --");
            println!("'{}'", output);
        }
        println!("result: {}", f.result());
        Ok(())
    }

    #[test]
    fn foo() {
    }
}