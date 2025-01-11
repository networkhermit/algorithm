use super::*;

use crate::algorithms::search::tests::{derive_empty, derive_increasing};

#[test]
fn test_linear_search() {
    derive_empty(&find)();
    derive_increasing(&find, 10000)();
}
