fn fizz_buzz(n: i32) -> String {
    let mut word = String::new();

    if n % 3 == 0 {
        word += "Fizz";
    }
    if n % 5 == 0 {
        word += "Buzz";
    }

    if word.is_empty() {
        word = n.to_string();
    }

    word
}

fn main() {
    (1..=100).for_each(|i| {
        println!("{}", fizz_buzz(i));
    });
}
