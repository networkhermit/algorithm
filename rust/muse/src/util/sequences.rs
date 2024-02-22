use super::random_factory;

pub fn inspect(arr: &[i32]) {
    println!("[");
    arr.iter().enumerate().for_each(|(i, v)| {
        println!("\t#{:#04X}  ->  {}", i, v);
    });
    println!("]");
}

pub fn is_sorted(arr: &[i32]) -> bool {
    for i in 1..arr.len() {
        if arr[i] < arr[i - 1] {
            return false;
        }
    }

    true
}

pub fn parity_checksum(arr: &[i32]) -> i32 {
    arr.iter().fold(0, |checksum, v| checksum ^ v)
}

pub fn reverse(arr: &mut [i32]) {
    (0..arr.len() >> 1).for_each(|i| {
        let k = arr.len() - i - 1;
        arr.swap(i, k);
    });
}

pub fn shuffle(arr: &mut [i32]) {
    (0..arr.len()).for_each(|i| {
        let k = random_factory::gen_int_n(i as i32, (arr.len() - 1) as i32);
        arr.swap(i, k as usize);
    });
}

pub fn sort(arr: &mut [i32]) {
    arr.sort_unstable();
}
