use super::{List, Resizable};

#[derive(Debug)]
pub struct ArrayList<T> {
    data: Vec<Option<T>>,
    logical_size: usize,
    physical_size: usize,
}

impl<T> ArrayList<T> {
    const DEFAULT_CAPACITY: usize = 64;

    pub fn new() -> Self {
        Self::new_with_capacity(0)
    }

    pub fn new_with_capacity(physical_size: usize) -> Self {
        let mut list = Self {
            data: Default::default(),
            logical_size: 0,
            physical_size: Self::DEFAULT_CAPACITY,
        };

        if physical_size > 1 {
            list.physical_size = physical_size;
        }
        list.data = Vec::with_capacity(list.physical_size);
        (0..list.physical_size).for_each(|_| {
            list.data.push(None);
        });

        list
    }
}

impl<T> Default for ArrayList<T> {
    fn default() -> Self {
        Self::new()
    }
}

impl<T> List<T> for ArrayList<T> {
    fn size(&self) -> usize {
        self.logical_size
    }

    fn is_empty(&self) -> bool {
        self.logical_size == 0
    }

    fn get(&self, index: usize) -> &T {
        if index >= self.logical_size {
            panic!("[PANIC - IndexOutOfBounds]");
        }

        self.data[index].as_ref().expect("element should exist")
    }

    fn set(&mut self, index: usize, element: T) {
        if index >= self.logical_size {
            panic!("[PANIC - IndexOutOfBounds]");
        }

        self.data[index] = Some(element);
    }

    fn insert(&mut self, index: usize, element: T) {
        if index > self.logical_size {
            panic!("[PANIC - IndexOutOfBounds]");
        }

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

        for i in (index + 1..=self.logical_size).rev() {
            self.data[i] = self.data[i - 1].take();
        }

        self.data[index] = Some(element);

        self.logical_size += 1;
    }

    fn remove(&mut self, index: usize) {
        if index >= self.logical_size {
            panic!("[PANIC - IndexOutOfBounds]");
        }

        for i in index + 1..self.logical_size {
            self.data[i - 1] = self.data[i].take();
        }

        self.logical_size -= 1;
    }

    fn front(&self) -> &T {
        self.get(0)
    }

    fn back(&self) -> &T {
        self.get(self.logical_size - 1)
    }

    fn prepend(&mut self, element: T) {
        self.insert(0, element)
    }

    fn append(&mut self, element: T) {
        self.insert(self.logical_size, element)
    }

    fn poll(&mut self) {
        self.remove(0)
    }

    fn eject(&mut self) {
        self.remove(self.logical_size - 1)
    }
}

impl<T> Resizable for ArrayList<T> {
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
