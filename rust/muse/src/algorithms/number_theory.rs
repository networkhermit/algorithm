pub mod coprimality;
pub mod factorial;
pub mod fibonacci_number;
pub mod greatest_common_divisor;
pub mod least_common_multiple;
pub mod parity;
pub mod primality;
pub mod prime_sieves;

pub fn gcd(m: i64, n: i64) -> i64 {
    greatest_common_divisor::iterative_binary_gcd(m, n)
}

pub fn is_even(n: i64) -> bool {
    parity::bitwise_is_even(n)
}

pub fn is_odd(n: i64) -> bool {
    parity::bitwise_is_odd(n)
}

#[cfg(test)]
pub(crate) mod tests;
