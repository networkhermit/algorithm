pub fn sort(arr: &mut [i32]) {
    if arr.is_empty() {
        return;
    }

    for i in 0..arr.len() - 1 {
        let mut i_min = i;
        for j in i + 1..arr.len() {
            if arr[j] < arr[i_min] {
                i_min = j;
            }
        }
        if i_min != i {
            (arr[i], arr[i_min]) = (arr[i_min], arr[i])
        }
    }
}

#[cfg(test)]
mod tests;
