pub fn iterative_procedure(n: i64) -> i64 {
    if n < 0 {
        return 0;
    }

    let mut result = 1;
    (1..=n).for_each(|i| {
        result *= i;
    });

    result
}

pub fn recursive_procedure(n: i64) -> i64 {
    if n < 0 {
        return 0;
    }

    if n == 0 {
        return 1;
    }
    recursive_procedure(n - 1) * n
}

#[cfg(test)]
pub(crate) mod tests;
