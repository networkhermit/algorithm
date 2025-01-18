use crate::data_structures::tests::{stack_derive, stack_derive_resizable};

use super::*;

#[test]
fn test_array_stack() {
    stack_derive::<ArrayStack<_>>()();
    stack_derive_resizable::<ArrayStack<_>>()();
}
