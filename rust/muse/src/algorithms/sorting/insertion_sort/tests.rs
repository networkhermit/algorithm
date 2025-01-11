use super::*;

use crate::algorithms::sorting::tests::{
    derive_decreasing, derive_empty, derive_identical, derive_increasing, derive_random,
};

#[test]
fn test_insertion_sort() {
    derive_decreasing(&sort, 10000)();
    derive_empty(&sort)();
    derive_identical(&sort, 100000)();
    derive_increasing(&sort, 100000)();
    derive_random(&sort, 10000)();
}
