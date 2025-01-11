pub mod bubble_sort;
pub mod insertion_sort;
pub mod merge_sort;
pub mod quick_sort;
pub mod selection_sort;

pub fn bubble_sort(arr: &mut [i32]) {
    bubble_sort::sort(arr);
}

pub fn insertion_sort(arr: &mut [i32]) {
    insertion_sort::sort(arr);
}

pub fn merge_sort(arr: &mut [i32]) {
    merge_sort::sort(arr);
}

pub fn quick_sort(arr: &mut [i32]) {
    quick_sort::sort(arr);
}

pub fn selection_sort(arr: &mut [i32]) {
    selection_sort::sort(arr);
}

#[cfg(test)]
pub(crate) mod tests;
