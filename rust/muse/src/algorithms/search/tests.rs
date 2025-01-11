use crate::util::sequence_builder;

use super::*;

pub(crate) fn derive<'a>(
    f: &'a dyn Fn(&[i32], i32) -> usize,
    size: usize,
    pack: &'a dyn Fn(&mut [i32]),
) -> impl Fn() + 'a {
    move || {
        let mut arr = vec![0; size];
        pack(&mut arr);

        let sentinel = [-1, 2_147_483_647];

        for v in sentinel {
            let actual = f(&arr, v);
            assert_eq!(
                actual, size,
                "fn(arr, {}) returned the left result, expect the right result",
                v
            );
        }

        for (i, &v) in arr.iter().enumerate() {
            let actual = f(&arr, v);
            assert_eq!(
                actual, i,
                "fn(arr, {}) returned the left result, expect the right result",
                v
            );
        }
    }
}

pub(crate) fn derive_empty(f: &dyn Fn(&[i32], i32) -> usize) -> impl Fn() + '_ {
    derive(f, 0, &sequence_builder::pack_identical)
}

pub(crate) fn derive_increasing(f: &dyn Fn(&[i32], i32) -> usize, size: usize) -> impl Fn() + '_ {
    derive(f, size, &sequence_builder::pack_increasing)
}

#[test]
fn test_binary_search() {
    derive_increasing(&binary_search, 100000)();
}

#[test]
fn test_linear_search() {
    derive_increasing(&linear_search, 10000)();
}
