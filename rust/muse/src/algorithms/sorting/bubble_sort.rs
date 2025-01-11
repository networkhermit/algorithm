pub fn sort(arr: &mut [i32]) {
    let mut unsorted = arr.len();
    while unsorted > 1 {
        let mut margin = 0;
        for i in 1..unsorted {
            if arr[i - 1] > arr[i] {
                (arr[i - 1], arr[i]) = (arr[i], arr[i - 1]);
                margin = i;
            }
        }
        unsorted = margin;
    }
}

#[cfg(test)]
mod tests;
