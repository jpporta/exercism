pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let mut res = Vec::new();
    let mut y = 0;
    let no_lines = minefield.len();
    for line in minefield {
        let mut new_line = String::new();
        let chars = line.as_bytes();
        let mut x = 0;
        let prev_line = match y {
            0 => None,
            _ => Some(minefield[y - 1].as_bytes()),
        };
        let next_line = match y >= no_lines - 1{
            true => None,
            _ => Some(minefield[y + 1].as_bytes()),
        };
        for c in chars {
            if *c as char == '*' {
                new_line.push('*');
                x+=1;
                continue;
            }
            let mut sum: u8 = 0;
            if let Some(l) = prev_line {
                if x > 0 && let Some(v) = l.get(x-1) && *v as char == '*' {
                    sum+=1;
                }
                if let Some(v) = l.get(x) && *v as char == '*' {
                    sum+=1;
                }
                if let Some(v) = l.get(x+1) && *v as char == '*' {
                    sum+=1;
                }
            }
            if let Some(l) = next_line {
                if x > 0 && let Some(v) = l.get(x-1) && *v as char == '*' {
                    sum+=1;
                }
                if let Some(v) = l.get(x) && *v as char == '*' {
                    sum+=1;
                }
                if let Some(v) = l.get(x+1) && *v as char == '*' {
                    sum+=1;
                }
            }
            if x > 0 && let Some(v) = chars.get(x-1) && *v as char == '*' {
                    sum+=1;
                }
            if let Some(v) = chars.get(x+1) && *v as char == '*' {
                    sum+=1;
                }
            println!("at {x},{y} -> sum: {sum}. Chars: {chars:?}");

            match sum.to_string().chars().next() {
                None => {}
                Some('0') => new_line.push(' '),
                Some(v) => new_line.push(v),
            }
            x += 1;
        }
        res.push(new_line);
        y += 1;
    }
    res
}
