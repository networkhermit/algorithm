pub mod binary_search;
pub mod linear_search;

pub fn binary_search(arr: &[i32], key: i32) -> usize {
    binary_search::find(arr, key)
}

pub fn linear_search(arr: &[i32], key: i32) -> usize {
    linear_search::find(arr, key)
}

#[cfg(test)]
pub(crate) mod tests;
