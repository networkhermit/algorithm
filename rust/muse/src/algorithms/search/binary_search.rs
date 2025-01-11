use std::cmp::Ordering;

pub fn find(arr: &[i32], key: i32) -> usize {
    let mut lo = 0;
    let mut hi = arr.len();

    while lo < hi {
        let mid = lo + ((hi - lo) >> 1);
        match key.cmp(&arr[mid]) {
            Ordering::Less => {
                hi = mid;
            }
            Ordering::Greater => {
                lo = mid + 1;
            }
            Ordering::Equal => {
                return mid;
            }
        };
    }

    arr.len()
}

#[cfg(test)]
mod tests;
