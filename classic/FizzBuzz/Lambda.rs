fn fizz_buzz(n: i32) -> String {
    let mut word = String::new();

    if n % 3 == 0 {
        word += "Fizz";
    }
    if n % 5 == 0 {
        word += "Buzz";
    }

    if word.len() == 0 {
        word = n.to_string();
    }

    word
}

fn main() {
    for i in 1..=100 {
        println!("{}", fizz_buzz(i));
    }
}
