use std::collections::HashMap;

type StdResult<T> = Result<T, Box<dyn std::error::Error>>;

fn read_field(input: &str) -> StdResult<Vec<Vec<u8>>> {
    let mut field = vec![];
    for line in input.lines() {
        field.push(line.split("").filter(|s| *s != "").map(|i| i.parse() ).collect::<Result<_,_>>()?);
    }
    Ok(field)
}

fn main() -> StdResult<()> {
    let input = std::fs::read_to_string("input")?;

    let field = read_field(input.as_str())?;

    let height = field.len();
    let width = field[0].len();

    let mut low_points: Vec<(usize, usize)> = vec![];
    for row in 0..height {
        for col in 0..width {
            if ( row == 0 || field[row][col] < field[row-1][col] ) &&
                ( col == width - 1 || field[row][col] < field[row][col + 1] ) &&
                ( row == height - 1 || field[row][col] < field[row+1][col] ) &&
                (col == 0 || field[row][col] < field[row][col - 1]) {
                    low_points.push((col, row))
            };
        }
    }
    let sum: u32 = low_points.iter().map(|(col, row)| (field[*row][*col] + 1) as u32 ).sum();
    println!("part one: low_points: {}", sum);

    let mut basins = vec![vec![0 as u32; width]; height];
    let mut bi: u32 = 1;
    for low_point in low_points {
        mark_basin(low_point, bi, &field, &mut &mut basins);
        bi += 1;
    }

    let mut basin_size = HashMap::new();
    for row in basins {
        for col in row {
            if col == 0 {
                continue
            }
            *basin_size.entry(col).or_insert(0) += 1;
        }
    }

    let mut basins_list: Vec<(&u32, &i32)> = basin_size.iter().collect();
    basins_list.sort_by(|a, b| b.1.cmp(a.1));
    let result: i32 = basins_list.iter().take(3).map(|(a, b)| **b).fold(1, |a, x| a * x);
    println!("part two: top three basins multiplied: {}", result);

    Ok(())
}

fn mark_basin(point: (usize, usize), num: u32, map: &Vec<Vec<u8>>, basins: &mut Vec<Vec<u32>>) {
    let (col, row) = point;
    let val = map[row][col];

    // height 9 does not belong to basin
    if val == 9 {
        return
    }

    // if not yet in basin (zero mean it does not yet belong to a basin)
    if basins[row][col] != 0 {
        return
    }
    basins[row][col] = num;

    let height = map.len();
    let width = map[0].len();

    if row != 0 && map[row-1][col] > val {
        mark_basin((col, row-1), num, map, basins)
    }
    if !(row >= height- 1 ) && map[row+1][col] > val {
        mark_basin((col, row+1), num, map, basins)
    }
    if col != 0 && map[row][col - 1] > val {
        mark_basin((col-1, row), num, map, basins)
    }
    if !(col >= width-1) && map[row][col + 1] > val {
        mark_basin((col+1, row), num, map, basins)
    }
}