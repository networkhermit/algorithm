use crate::data_structures::tests::{queue_derive, queue_derive_resizable};

use super::*;

#[test]
fn test_array_queue() {
    queue_derive::<ArrayQueue<_>>()();
    queue_derive_resizable::<ArrayQueue<_>>()();
}
