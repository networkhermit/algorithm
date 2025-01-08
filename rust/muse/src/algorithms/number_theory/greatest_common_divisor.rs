pub fn iterative_binary_gcd(mut m: i64, mut n: i64) -> i64 {
    if m < 0 {
        m = -m;
    }
    if n < 0 {
        n = -n;
    }

    let mut shift = 0;

    loop {
        if m == n {
            return m << shift;
        }
        if m == 0 {
            return n << shift;
        }
        if n == 0 {
            return m << shift;
        }

        if (m & 1) == 0 {
            m >>= 1;
            if (n & 1) == 0 {
                n >>= 1;
                shift += 1;
            }
        } else if (n & 1) == 0 {
            n >>= 1;
        } else if m > n {
            m = (m - n) >> 1;
        } else {
            n = (n - m) >> 1;
        }
    }
}

pub fn recursive_binary_gcd(mut m: i64, mut n: i64) -> i64 {
    if m < 0 {
        m = -m;
    }
    if n < 0 {
        n = -n;
    }

    if m == n {
        return m;
    }
    if m == 0 {
        return n;
    }
    if n == 0 {
        return m;
    }

    if (m & 1) == 0 {
        if (n & 1) == 0 {
            return recursive_binary_gcd(m >> 1, n >> 1) << 1;
        }
        return recursive_binary_gcd(m >> 1, n);
    }
    if (n & 1) == 0 {
        return recursive_binary_gcd(m, n >> 1);
    }
    if m > n {
        return recursive_binary_gcd((m - n) >> 1, n);
    }
    recursive_binary_gcd(m, (n - m) >> 1)
}

pub fn iterative_euclidean(mut m: i64, mut n: i64) -> i64 {
    if m < 0 {
        m = -m;
    }
    if n < 0 {
        n = -n;
    }

    while n != 0 {
        (m, n) = (n, m % n);
    }

    m
}

pub fn recursive_euclidean(mut m: i64, mut n: i64) -> i64 {
    if m < 0 {
        m = -m;
    }
    if n < 0 {
        n = -n;
    }

    if n == 0 {
        return m;
    }
    recursive_euclidean(n, m % n)
}

#[cfg(test)]
pub(crate) mod tests;
