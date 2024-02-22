use super::{random_factory, sequences};

pub fn pack_identical(arr: &mut [i32]) {
    let n = random_factory::gen_int();
    arr.fill(n);
}

pub fn pack_increasing(arr: &mut [i32]) {
    if arr.is_empty() {
        return;
    }
    arr[0] = random_factory::gen_int_n(1, 3);
    (1..arr.len()).for_each(|i| {
        arr[i] = arr[i - 1] + random_factory::gen_int_n(1, 3);
    });
}

pub fn pack_random(arr: &mut [i32]) {
    arr.fill_with(random_factory::gen_int)
}

pub fn pack_decreasing(arr: &mut [i32]) {
    pack_increasing(arr);
    sequences::reverse(arr);
}
