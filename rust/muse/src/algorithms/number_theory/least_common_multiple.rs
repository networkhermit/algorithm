use super::greatest_common_divisor;

pub fn reduce_to_binary_gcd(mut m: i64, mut n: i64) -> i64 {
    if m < 0 {
        m = -m;
    }
    if n < 0 {
        n = -n;
    }

    if m == 0 || n == 0 {
        return 0;
    }
    m / greatest_common_divisor::iterative_binary_gcd(m, n) * n
}

pub fn reduce_to_euclidean(mut m: i64, mut n: i64) -> i64 {
    if m < 0 {
        m = -m;
    }
    if n < 0 {
        n = -n;
    }

    if m == 0 || n == 0 {
        return 0;
    }

    m / greatest_common_divisor::iterative_euclidean(m, n) * n
}

#[cfg(test)]
pub(crate) mod tests;
