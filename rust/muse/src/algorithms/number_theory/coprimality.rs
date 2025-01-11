use super::greatest_common_divisor;

pub fn reduce_to_binary_gcd(m: i64, n: i64) -> bool {
    greatest_common_divisor::iterative_binary_gcd(m, n) == 1
}

pub fn reduce_to_euclidean(m: i64, n: i64) -> bool {
    greatest_common_divisor::iterative_euclidean(m, n) == 1
}

#[cfg(test)]
pub(crate) mod tests;
