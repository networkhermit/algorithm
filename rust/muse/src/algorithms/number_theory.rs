pub mod coprimality;
pub mod factorial;
pub mod fibonacci_number;
pub mod greatest_common_divisor;
pub mod least_common_multiple;
pub mod parity;
pub mod primality;
pub mod prime_sieves;

pub fn is_coprime(m: i64, n: i64) -> bool {
    coprimality::reduce_to_binary_gcd(m, n)
}

pub fn factorial(n: i64) -> i64 {
    factorial::iterative_procedure(n)
}

pub fn fibonacci(n: i64) -> i64 {
    fibonacci_number::iterative_procedure(n)
}

pub fn gcd(m: i64, n: i64) -> i64 {
    greatest_common_divisor::iterative_binary_gcd(m, n)
}

pub fn lcm(m: i64, n: i64) -> i64 {
    least_common_multiple::reduce_to_binary_gcd(m, n)
}

pub fn is_even(n: i64) -> bool {
    parity::bitwise_is_even(n)
}

pub fn is_odd(n: i64) -> bool {
    parity::bitwise_is_odd(n)
}

pub fn is_prime(n: i64) -> bool {
    primality::is_prime(n)
}

pub fn is_composite(n: i64) -> bool {
    primality::is_composite(n)
}

pub fn sieve_of_primes(n: i32) -> Vec<i32> {
    prime_sieves::sieve_of_eratosthenes(n)
}

#[cfg(test)]
pub(crate) mod tests;
