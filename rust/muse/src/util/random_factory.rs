use rand::{thread_rng, Rng};

pub fn gen_int_n(min: i32, max: i32) -> i32 {
    let mut rng = thread_rng();
    rng.gen_range(min..=max)
}

pub fn gen_int() -> i32 {
    gen_int_n(1, 2_147_483_647)
}

pub fn gen_even() -> i32 {
    gen_int() >> 1 << 1
}

pub fn gen_odd() -> i32 {
    gen_int() | 1
}

#[cfg(test)]
mod tests;
