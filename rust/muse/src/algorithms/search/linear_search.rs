pub fn find(arr: &[i32], key: i32) -> usize {
    for (i, &v) in arr.iter().enumerate() {
        if v == key {
            return i;
        }
    }

    arr.len()
}

#[cfg(test)]
mod tests;
