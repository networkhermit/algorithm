use super::*;

use crate::algorithms::selection::tests::{
    derive_decreasing, derive_empty, derive_identical, derive_increasing, derive_random,
};

#[test]
fn test_quick_select() {
    derive_decreasing(&find_kth_largest, 100000)();
    derive_empty(&find_kth_largest)();
    derive_identical(&find_kth_largest, 100000)();
    derive_increasing(&find_kth_largest, 100000)();
    derive_random(&find_kth_largest, 100000)();
}

#[test]
fn test_quick_select_three_way_partition() {
    derive_decreasing(&find_kth_largest_three_way_partition, 10000)();
    derive_empty(&find_kth_largest_three_way_partition)();
    derive_identical(&find_kth_largest_three_way_partition, 100000)();
    derive_increasing(&find_kth_largest_three_way_partition, 10000)();
    derive_random(&find_kth_largest_three_way_partition, 100000)();
    derive_decreasing(&find_kth_largest_three_way_partition_with_shuffle, 100000)();
    derive_increasing(&find_kth_largest_three_way_partition_with_shuffle, 100000)();
}

#[test]
fn test_quick_select_inefficient() {
    derive_decreasing(&find_kth_largest_inefficient, 10000)();
    derive_empty(&find_kth_largest_inefficient)();
    derive_identical(&find_kth_largest_inefficient, 10000)();
    derive_increasing(&find_kth_largest_inefficient, 10000)();
    derive_random(&find_kth_largest_inefficient, 100000)();
}
