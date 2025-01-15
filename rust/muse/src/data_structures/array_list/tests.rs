use crate::data_structures::tests::{list_derive, list_derive_resizable};

use super::*;

#[test]
fn test_array_list() {
    list_derive::<ArrayList<_>>()();
    list_derive_resizable::<ArrayList<_>>()();
}
