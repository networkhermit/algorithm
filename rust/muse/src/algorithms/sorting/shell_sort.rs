pub fn sort(arr: &mut [i32]) {
    let gaps = [701, 301, 132, 57, 23, 10, 4, 1];
    for gap in gaps.into_iter() {
        for i in gap..arr.len() {
            let target = arr[i];
            let mut cursor = i;
            while cursor >= gap {
                if arr[cursor - gap] <= target {
                    break;
                }
                arr[cursor] = arr[cursor - gap];
                cursor -= gap;
            }
            arr[cursor] = target;
        }
    }
}

#[cfg(test)]
mod tests;
