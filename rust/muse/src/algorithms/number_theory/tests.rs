use std::fmt::Display;

use super::*;

#[derive(Debug)]
pub(crate) struct UniqueCategorySample<T, C>(pub(crate) T, pub(crate) C);

pub(crate) fn unique_category_derive<'a, T: Copy + Display, C: Eq>(
    f: &'a dyn Fn(T) -> bool,
    sample: &'a [UniqueCategorySample<T, C>],
    c: C,
) -> impl Fn() + 'a {
    move || {
        sample.iter().for_each(|tc| {
            let actual = f(tc.0);
            let expected = tc.1 == c;
            assert_eq!(actual, expected, "fn({}) returned the left result", tc.0);
        });
    }
}

#[derive(Debug)]
pub(crate) struct MNUniqueCategorySample<T, C>(pub(crate) T, pub(crate) T, pub(crate) C);

pub(crate) fn mn_unique_category_derive<'a, T: Copy + Display, C: Eq>(
    f: &'a dyn Fn(T, T) -> bool,
    sample: &'a [MNUniqueCategorySample<T, C>],
    c: C,
) -> impl Fn() + 'a {
    move || {
        sample.iter().for_each(|tc| {
            let actual = f(tc.0, tc.1);
            let expected = tc.2 == c;
            assert_eq!(
                actual, expected,
                "fn({}, {}) returned the left result",
                tc.0, tc.1
            );
        });
    }
}

#[test]
fn test_is_coprime() {
    coprimality::tests::derive(&is_coprime)();
}

#[test]
fn test_factorial() {
    factorial::tests::derive(&factorial)();
}

#[test]
fn test_fibonacci() {
    fibonacci_number::tests::derive(&fibonacci)();
}

#[test]
fn test_gcd() {
    greatest_common_divisor::tests::derive(&gcd)();
}

#[test]
fn test_lcm() {
    least_common_multiple::tests::derive(&lcm)();
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

#[test]
fn test_is_prime() {
    use primality::tests::{derive, Category::Prime};

    derive(&is_prime, Prime)();
}

#[test]
fn test_is_composite() {
    use primality::tests::{derive, Category::Composite};

    derive(&is_composite, Composite)();
}

#[test]
fn test_sieve_of_primes() {
    use prime_sieves::tests::derive;

    derive(&sieve_of_primes, &primality::is_prime)();
}
