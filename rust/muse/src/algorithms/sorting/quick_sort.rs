use rand::{seq::SliceRandom, thread_rng};

pub fn partition(arr: &mut [i32], lo: usize, hi: usize) -> usize {
    let pivot = arr[lo + ((hi - 1 - lo) >> 1)];

    let mut i = lo.wrapping_sub(1);
    let mut j = hi;

    loop {
        loop {
            i = i.wrapping_add(1);
            if arr[i] >= pivot {
                break;
            }
        }
        loop {
            j -= 1;
            if arr[j] <= pivot {
                break;
            }
        }
        if i >= j {
            return j;
        }
        arr.swap(i, j);
    }
}

pub fn sort(arr: &mut [i32]) {
    sort_with_range(arr, 0, arr.len())
}

pub fn sort_with_range(arr: &mut [i32], lo: usize, hi: usize) {
    if lo + 1 >= hi {
        return;
    }

    let p = partition(arr, lo, hi);

    sort_with_range(arr, lo, p + 1);
    sort_with_range(arr, p + 1, hi);
}

pub fn partition_three_way(arr: &mut [i32], lo: usize, hi: usize) -> (usize, usize) {
    let mut i = lo.wrapping_sub(1);
    let mut j = hi - 1;

    let mut p = lo.wrapping_sub(1);
    let mut q = hi - 1;

    let pivot = arr[hi - 1];

    loop {
        loop {
            i = i.wrapping_add(1);
            if arr[i] >= pivot {
                break;
            }
        }
        loop {
            j -= 1;
            if pivot >= arr[j] {
                break;
            }
            if j == lo {
                break;
            }
        }
        if i >= j {
            break;
        }
        arr.swap(i, j);
        if arr[i] == pivot {
            p = p.wrapping_add(1);
            arr.swap(p, i);
        }
        if pivot == arr[j] {
            q -= 1;
            arr.swap(j, q);
        }
    }

    arr.swap(i, hi - 1);

    j = i.wrapping_sub(1);
    i += 1;

    let mut k = lo;
    while k < p && p != usize::MAX {
        arr.swap(k, j);
        j -= 1;
        k += 1;
    }
    k = hi - 2;
    while k > q {
        arr.swap(i, k);
        i += 1;
        k -= 1;
    }

    (i, j)
}

pub fn sort_three_way_partition_with_shuffle(arr: &mut [i32]) {
    let mut rng = thread_rng();
    arr.shuffle(&mut rng);
    sort_with_range_three_way_partition(arr, 0, arr.len())
}

pub fn sort_three_way_partition(arr: &mut [i32]) {
    sort_with_range_three_way_partition(arr, 0, arr.len())
}

pub fn sort_with_range_three_way_partition(arr: &mut [i32], lo: usize, hi: usize) {
    if lo + 1 >= hi {
        return;
    }

    let (i, j) = partition_three_way(arr, lo, hi);

    sort_with_range_three_way_partition(arr, lo, j.wrapping_add(1));
    sort_with_range_three_way_partition(arr, i, hi);
}

pub fn partition_inefficient(arr: &mut [i32], lo: usize, hi: usize) -> usize {
    let pivot = arr[lo];

    let mut i = lo;
    let mut j = hi - 1;

    while i != j {
        while j > i {
            if arr[j] < pivot {
                (arr[i], arr[j]) = (arr[j], pivot);
                break;
            }
            j -= 1;
        }
        while i < j {
            if arr[i] > pivot {
                (arr[j], arr[i]) = (arr[i], pivot);
                break;
            }
            i += 1;
        }
    }

    i
}

pub fn sort_inefficient(arr: &mut [i32]) {
    sort_with_range_inefficient(arr, 0, arr.len())
}

pub fn sort_with_range_inefficient(arr: &mut [i32], lo: usize, hi: usize) {
    if lo >= hi {
        return;
    }

    let p = partition_inefficient(arr, lo, hi);

    sort_with_range_inefficient(arr, lo, p);
    sort_with_range_inefficient(arr, p + 1, hi);
}

#[cfg(test)]
mod tests;
