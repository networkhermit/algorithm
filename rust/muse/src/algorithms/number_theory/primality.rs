pub fn is_prime(n: i64) -> bool {
    if n < 2 {
        return false;
    }
    if (n & 1) == 0 && n != 2 {
        return false;
    }

    let bound = n.isqrt() + 1;
    for i in (3..bound).step_by(2) {
        if n % i == 0 {
            return false;
        }
    }

    true
}

pub fn is_composite(n: i64) -> bool {
    n > 1 && !is_prime(n)
}

#[cfg(test)]
pub(crate) mod tests;
