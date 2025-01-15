use crate::data_structures::tests::list_derive;

use super::*;

#[test]
fn test_linked_list() {
    list_derive::<LinkedList<_>>()();
}
