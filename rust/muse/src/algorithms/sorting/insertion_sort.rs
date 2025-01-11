pub fn sort(arr: &mut [i32]) {
    for i in 1..arr.len() {
        let target = arr[i];
        let mut cursor = i;
        while cursor > 0 {
            if arr[cursor - 1] <= target {
                break;
            }
            arr[cursor] = arr[cursor - 1];
            cursor -= 1;
        }
        arr[cursor] = target;
    }
}

#[cfg(test)]
mod tests;
