pub fn iterative_procedure(mut n: i64) -> i64 {
    let mut sign = 1;

    if n < 0 {
        if (n & 1) == 0 {
            sign = -1;
        }
        n = -n;
    }

    if n < 2 {
        return n;
    }

    let mut prev = 0;
    let mut curr = 1;

    for _ in 2..=n {
        (prev, curr) = (curr, prev + curr);
    }

    sign * curr
}

pub fn recursive_procedure(n: i64) -> i64 {
    if n < 0 {
        if (n & 1) == 0 {
            return -recursive_procedure(-n);
        }
        return recursive_procedure(-n);
    }
    if n < 2 {
        return n;
    }
    recursive_procedure(n - 2) + recursive_procedure(n - 1)
}

#[cfg(test)]
pub(crate) mod tests;
