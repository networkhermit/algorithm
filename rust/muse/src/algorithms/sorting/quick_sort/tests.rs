use super::*;

use crate::algorithms::sorting::tests::{
    derive_decreasing, derive_empty, derive_identical, derive_increasing, derive_random,
};

#[test]
fn test_quick_sort() {
    derive_decreasing(&sort, 100000)();
    derive_empty(&sort)();
    derive_identical(&sort, 100000)();
    derive_increasing(&sort, 100000)();
    derive_random(&sort, 100000)();
}

#[test]
fn test_quick_sort_three_way_partition() {
    derive_decreasing(&sort_three_way_partition, 10000)();
    derive_empty(&sort_three_way_partition)();
    derive_identical(&sort_three_way_partition, 100000)();
    derive_increasing(&sort_three_way_partition, 10000)();
    derive_random(&sort_three_way_partition, 100000)();
    derive_decreasing(&sort_three_way_partition_with_shuffle, 100000)();
    derive_increasing(&sort_three_way_partition_with_shuffle, 100000)();
}

#[test]
fn test_quick_sort_inefficient() {
    derive_decreasing(&sort_inefficient, 10000)();
    derive_empty(&sort_inefficient)();
    derive_identical(&sort_inefficient, 10000)();
    derive_increasing(&sort_inefficient, 10000)();
    derive_random(&sort_inefficient, 100000)();
}
