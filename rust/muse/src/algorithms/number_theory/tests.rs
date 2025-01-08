use std::fmt::Display;

use super::*;

#[derive(Debug)]
pub(crate) struct UniqueCategorySample<N, C>(pub(crate) N, pub(crate) C);

pub(crate) fn unique_category_derive<'a, N: Copy + Display, C: Eq>(
    f: &'a dyn Fn(N) -> bool,
    sample: &'a [UniqueCategorySample<N, C>],
    c: C,
) -> impl Fn() + 'a {
    move || {
        sample.iter().for_each(|tt| {
            let actual = f(tt.0);
            let expected = tt.1 == c;
            assert_eq!(actual, expected, "fn({}) returned the left result", tt.0);
        });
    }
}

#[test]
fn test_gcd() {
    greatest_common_divisor::tests::derive(&gcd)();
}

#[test]
fn test_is_even() {
    use parity::tests::{derive, Category::Even};

    derive(&is_even, Even)();
}

#[test]
fn test_is_odd() {
    use parity::tests::{derive, Category::Odd};

    derive(&is_odd, Odd)();
}
