use super::{Resizable, Stack};

#[derive(Debug)]
pub struct ArrayStack<T> {
    data: Vec<Option<T>>,
    logical_size: usize,
    physical_size: usize,
}

impl<T> ArrayStack<T> {
    const DEFAULT_CAPACITY: usize = 64;

    pub fn new() -> Self {
        Self::new_with_capacity(0)
    }

    pub fn new_with_capacity(physical_size: usize) -> Self {
        let mut stack = Self {
            data: Default::default(),
            logical_size: 0,
            physical_size: Self::DEFAULT_CAPACITY,
        };

        if physical_size > 1 {
            stack.physical_size = physical_size;
        }
        stack.data = Vec::with_capacity(stack.physical_size);
        (0..stack.physical_size).for_each(|_| {
            stack.data.push(None);
        });

        stack
    }
}

impl<T> Default for ArrayStack<T> {
    fn default() -> Self {
        Self::new()
    }
}

impl<T> Stack<T> for ArrayStack<T> {
    fn size(&self) -> usize {
        self.logical_size
    }

    fn is_empty(&self) -> bool {
        self.logical_size == 0
    }

    fn peek(&self) -> &T {
        if self.logical_size == 0 {
            panic!("[PANIC - NoSuchElement]");
        }

        self.data[self.logical_size - 1]
            .as_ref()
            .expect("element should exist")
    }

    fn push(&mut self, element: T) {
        if self.logical_size == self.physical_size {
            let mut new_capacity = Self::DEFAULT_CAPACITY;

            if self.physical_size >= Self::DEFAULT_CAPACITY {
                new_capacity = self.physical_size + (self.physical_size >> 1)
            }

            let mut temp = Vec::with_capacity(new_capacity);

            (0..self.logical_size).for_each(|i| temp.push(self.data[i].take()));
            (self.logical_size..new_capacity).for_each(|_| {
                temp.push(None);
            });

            self.data = temp;
            self.physical_size = new_capacity;
        }

        self.data[self.logical_size] = Some(element);

        self.logical_size += 1;
    }

    fn pop(&mut self) {
        if self.logical_size == 0 {
            panic!("[PANIC - NoSuchElement]");
        }

        self.logical_size -= 1;

        self.data[self.logical_size].take();
    }
}

impl<T> Resizable for ArrayStack<T> {
    fn capacity(&self) -> usize {
        self.physical_size
    }

    fn shrink(&mut self) {
        let mut temp = Vec::with_capacity(self.logical_size);

        (0..self.logical_size).for_each(|i| {
            temp.push(self.data[i].take());
        });

        self.data = temp;
        self.physical_size = self.logical_size;
    }
}

#[cfg(test)]
pub(crate) mod tests;
