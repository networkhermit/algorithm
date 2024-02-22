pub fn modulo_is_even(n: i64) -> bool {
    n % 2 == 0
}

pub fn modulo_is_odd(n: i64) -> bool {
    n % 2 != 0
}

pub fn bitwise_is_even(n: i64) -> bool {
    (n & 1) == 0
}

pub fn bitwise_is_odd(n: i64) -> bool {
    (n & 1) != 0
}

#[cfg(test)]
pub(crate) mod tests;
