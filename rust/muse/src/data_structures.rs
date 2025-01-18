pub mod array_list;
pub mod array_queue;
pub mod array_stack;
pub mod linked_list;
pub mod singly_linked_list;

pub trait List<T> {
    fn size(&self) -> usize;

    fn is_empty(&self) -> bool;

    fn get(&self, index: usize) -> &T;

    fn set(&mut self, index: usize, element: T);

    fn insert(&mut self, index: usize, element: T);

    fn remove(&mut self, index: usize);

    fn front(&self) -> &T;

    fn back(&self) -> &T;

    fn prepend(&mut self, element: T);

    fn append(&mut self, element: T);

    fn poll(&mut self);

    fn eject(&mut self);
}

pub trait Queue<T> {
    fn size(&self) -> usize;

    fn is_empty(&self) -> bool;

    fn peek(&self) -> &T;

    fn enqueue(&mut self, element: T);

    fn dequeue(&mut self);
}

pub trait Stack<T> {
    fn size(&self) -> usize;

    fn is_empty(&self) -> bool;

    fn peek(&self) -> &T;

    fn push(&mut self, element: T);

    fn pop(&mut self);
}

pub trait Resizable {
    fn capacity(&self) -> usize;

    fn shrink(&mut self);
}

#[cfg(test)]
pub(crate) mod tests;
