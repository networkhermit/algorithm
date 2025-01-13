pub fn sieve_of_eratosthenes(n: i32) -> Vec<i32> {
    if n < 2 {
        return vec![];
    }

    let size = (n as usize + 1) >> 1;

    let mut arr = vec![false; size];

    let mut num_primes = size;
    let bound = (n as usize).isqrt() + 1;
    for i in (3..bound).step_by(2) {
        if arr[i >> 1] {
            continue;
        }
        let mut j = i * i;
        while j <= n as usize {
            if !arr[j >> 1] {
                arr[j >> 1] = true;
                num_primes -= 1;
            }
            j += i << 1;
        }
    }

    let mut primes = Vec::with_capacity(num_primes);

    primes.push(2);

    for i in (3..=n as usize).step_by(2) {
        if !arr[i >> 1] {
            primes.push(i as i32);
        }
    }

    primes
}

#[cfg(test)]
pub(crate) mod tests;
