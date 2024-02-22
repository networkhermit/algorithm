use rand::{seq::SliceRandom, thread_rng};

use crate::algorithms::sorting::quick_sort::{
    partition, partition_inefficient, partition_three_way,
};

pub fn find_kth_largest(arr: &mut [i32], k: usize) -> i32 {
    let length = arr.len();
    quick_select(arr, 0, length, length - k).unwrap_or_default()
}

fn quick_select(arr: &mut [i32], lo: usize, hi: usize, k: usize) -> Option<i32> {
    if lo + 1 == hi {
        return match lo == k {
            true => Some(arr[lo]),
            false => None,
        };
    }

    let p = partition(arr, lo, hi);

    match p {
        p if k <= p => quick_select(arr, lo, p + 1, k),
        _ => quick_select(arr, p + 1, hi, k),
    }
}

pub fn find_kth_largest_three_way_partition_with_shuffle(arr: &mut [i32], k: usize) -> i32 {
    let length = arr.len();
    let mut rng = thread_rng();
    arr.shuffle(&mut rng);
    quick_select_three_way_partition(arr, 0, length, length - k).unwrap_or_default()
}

pub fn find_kth_largest_three_way_partition(arr: &mut [i32], k: usize) -> i32 {
    let length = arr.len();
    quick_select_three_way_partition(arr, 0, length, length - k).unwrap_or_default()
}

fn quick_select_three_way_partition(
    arr: &mut [i32],
    lo: usize,
    hi: usize,
    k: usize,
) -> Option<i32> {
    if lo + 1 >= hi {
        return arr.get(k).copied();
    }

    let (i, j) = partition_three_way(arr, lo, hi);

    match k {
        k if k <= j && j != usize::MAX => quick_select_three_way_partition(arr, lo, j + 1, k),
        k if k >= i => quick_select_three_way_partition(arr, i, hi, k),
        _ => arr.get(k).copied(),
    }
}

pub fn quick_select_three_way_partition_iterative(
    arr: &mut [i32],
    lo: usize,
    hi: usize,
    k: usize,
) -> Option<i32> {
    let mut lo = lo;
    let mut hi = hi;

    while lo + 1 < hi {
        let (i, j) = partition_three_way(arr, lo, hi);

        match k {
            k if k <= j && j != usize::MAX => {
                hi = j + 1;
            }
            k if k >= i => {
                lo = i;
            }
            _ => return arr.get(k).copied(),
        };
    }

    arr.get(k).copied()
}

pub fn find_kth_largest_inefficient(arr: &mut [i32], k: usize) -> i32 {
    let length = arr.len();
    quick_select_inefficient(arr, 0, length, length - k).unwrap_or_default()
}

fn quick_select_inefficient(arr: &mut [i32], lo: usize, hi: usize, k: usize) -> Option<i32> {
    if lo >= hi {
        return None;
    }

    let p = partition_inefficient(arr, lo, hi);

    match p {
        p if k < p => quick_select_inefficient(arr, lo, p, k),
        p if p == k => arr.get(k).copied(),
        _ => quick_select_inefficient(arr, p + 1, hi, k),
    }
}

#[cfg(test)]
mod tests;
