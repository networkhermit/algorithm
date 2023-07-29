pub mod parity;

pub fn is_even(n: i64) -> bool {
    parity::bitwise_is_even(n)
}

pub fn is_odd(n: i64) -> bool {
    parity::bitwise_is_odd(n)
}

#[cfg(test)]
mod tests {
    use std::fmt::Display;

    use super::*;

    #[derive(Debug)]
    pub(crate) struct UniqueCategorySample<N, C>(pub(crate) N, pub(crate) C);

    pub(crate) fn unique_category_derive<N: Copy + Display, C: Eq>(
        f: &dyn Fn(N) -> bool,
        sample: &[UniqueCategorySample<N, C>],
        c: C,
    ) {
        for tt in sample.iter() {
            let actual = f(tt.0);
            let expected = tt.1 == c;
            assert_eq!(actual, expected, "fn({}) returned the left result", tt.0);
        }
    }

    #[test]
    fn test_is_even() {
        use parity::tests::{derive, Category::Even};

        derive(&is_even, Even);
    }

    #[test]
    fn test_is_odd() {
        use parity::tests::{derive, Category::Odd};

        derive(&is_odd, Odd);
    }
}
