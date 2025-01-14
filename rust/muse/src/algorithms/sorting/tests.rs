use crate::util::{sequence_builder, sequences};

use super::*;

pub(crate) fn derive<'a>(
    sort: &'a dyn Fn(&mut [i32]),
    size: usize,
    pack: &'a dyn Fn(&mut [i32]),
) -> impl Fn() + 'a {
    move || {
        let mut arr = vec![0; size];
        pack(&mut arr);

        let checksum = sequences::parity_checksum(&arr);

        sort(&mut arr);

        assert_eq!(
            sequences::parity_checksum(&arr),
            checksum,
            "elements returned have been altered"
        );

        assert!(
            sequences::is_sorted(&arr),
            "elements returned are not in sorted order"
        );
    }
}

pub(crate) fn derive_decreasing(sort: &dyn Fn(&mut [i32]), size: usize) -> impl Fn() + '_ {
    derive(sort, size, &sequence_builder::pack_decreasing)
}

pub(crate) fn derive_empty(sort: &dyn Fn(&mut [i32])) -> impl Fn() + '_ {
    derive(sort, 0, &sequence_builder::pack_identical)
}

pub(crate) fn derive_identical(sort: &dyn Fn(&mut [i32]), size: usize) -> impl Fn() + '_ {
    derive(sort, size, &sequence_builder::pack_identical)
}

pub(crate) fn derive_increasing(sort: &dyn Fn(&mut [i32]), size: usize) -> impl Fn() + '_ {
    derive(sort, size, &sequence_builder::pack_increasing)
}

pub(crate) fn derive_random(sort: &dyn Fn(&mut [i32]), size: usize) -> impl Fn() + '_ {
    derive(sort, size, &sequence_builder::pack_random)
}

#[test]
fn test_bubble_sort() {
    derive_random(&bubble_sort, 10000)();
}

#[test]
fn test_insertion_sort() {
    derive_random(&insertion_sort, 10000)();
}

#[test]
fn test_merge_sort() {
    derive_random(&merge_sort, 100000)();
}

#[test]
fn test_quick_sort() {
    derive_random(&quick_sort, 100000)();
}

#[test]
fn test_selection_sort() {
    derive_random(&selection_sort, 10000)();
}

#[test]
fn test_shell_sort() {
    derive_random(&shell_sort, 100000)();
}
