pub mod parity;

pub fn is_even(n: i64) -> bool {
    parity::bitwise_is_even(n)
}

pub fn is_odd(n: i64) -> bool {
    parity::bitwise_is_odd(n)
}

#[cfg(test)]
pub(crate) mod tests;
