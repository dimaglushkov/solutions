use std::fs::File;
use std::io::{self, BufRead};
use std::io::BufReader;
use std::path::Path;

fn read_lines<P>(filename: P) -> io::Result<Vec<String>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    let reader = BufReader::new(file);
    reader.lines().collect()
}

fn find_digits(s: &str) -> Vec<char> {
    s.chars().filter(|c| c.is_digit(10)).collect()
}

fn find_digits_and_words(s: &str) -> Vec<char> {
    let mut result = Vec::new();

    let word_to_digit = |word: &str| -> Option<char> {
        return if word.ends_with("one") {
            Some('1')
        } else if word.ends_with("two") {
            Some('2')
        } else if word.ends_with("three") {
            Some('3')
        } else if word.ends_with("four") {
            Some('4')
        } else if word.ends_with("five") {
            Some('5')
        } else if word.ends_with("six") {
            Some('6')
        } else if word.ends_with("seven") {
            Some('7')
        } else if word.ends_with("eight") {
            Some('8')
        } else if word.ends_with("nine") {
            Some('9')
        } else {
            None
        }
    };

    let mut word = String::new();
    for c in s.chars() {
        if c.is_alphabetic() {
            word.push(c);
            if let Some(digit) = word_to_digit(&word) {
                result.push(digit);
            }
        } else if c.is_digit(10) {
            result.push(c);
        }
    }

    result
}


fn main() {
    // let filename = "day01.example.input";
    let filename = "day01.input";
    let mut total = 0;
    match read_lines(filename) {
        Ok(lines) => {
            for line in &lines {
                let digits = find_digits(&line);
                if let (Some(first_digit), Some(last_digit)) = (digits.first(), digits.last()) {
                    let first_digit_val = first_digit.to_digit(10).unwrap_or(0) as i32;
                    let last_digit_val = last_digit.to_digit(10).unwrap_or(0) as i32;
                    total += first_digit_val * 10 + last_digit_val;
                }
            }
            println!("{}", total);

            total = 0;
            for line in &lines {
                let digits = find_digits_and_words(&line);
                if let (Some(first_digit), Some(last_digit)) = (digits.first(), digits.last()) {
                    let first_digit_val = first_digit.to_digit(10).unwrap_or(0) as i32;
                    let last_digit_val = last_digit.to_digit(10).unwrap_or(0) as i32;
                    total += first_digit_val * 10 + last_digit_val;
                }
            }
            println!("{}", total);
        }
        Err(e) => {
            println!("Failed to read from file: {}", e);
        }
    }


}