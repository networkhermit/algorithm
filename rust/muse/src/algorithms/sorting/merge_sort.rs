pub fn merge(arr: &mut [i32], lo: usize, mid: usize, hi: usize) {
    if lo == mid {
        return;
    }

    merge(arr, lo, (lo + mid) >> 1, mid);
    merge(arr, mid, (mid + hi) >> 1, hi);

    let mut m = lo;
    let mut n = mid;

    let mut sorted = vec![0; hi - lo];

    (0..sorted.len()).for_each(|i| {
        if m != mid && (n == hi || arr[m] < arr[n]) {
            sorted[i] = arr[m];
            m += 1;
        } else {
            sorted[i] = arr[n];
            n += 1;
        }
    });

    let mut cursor = 0;
    (lo..hi).for_each(|i| {
        arr[i] = sorted[cursor];
        cursor += 1;
    });
}

pub fn sort(arr: &mut [i32]) {
    merge(arr, 0, arr.len() >> 1, arr.len());
}

#[cfg(test)]
mod tests;
