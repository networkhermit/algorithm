use super::{Queue, Resizable};

#[derive(Debug)]
pub struct ArrayQueue<T> {
    data: Vec<Option<T>>,
    front: usize,
    logical_size: usize,
    physical_size: usize,
}

impl<T> ArrayQueue<T> {
    const DEFAULT_CAPACITY: usize = 64;

    pub fn new() -> Self {
        Self::new_with_capacity(0)
    }

    pub fn new_with_capacity(physical_size: usize) -> Self {
        let mut queue = Self {
            data: Default::default(),
            front: 0,
            logical_size: 0,
            physical_size: Self::DEFAULT_CAPACITY,
        };

        if physical_size > 1 {
            queue.physical_size = physical_size;
        }
        queue.data = Vec::with_capacity(queue.physical_size);
        (0..queue.physical_size).for_each(|_| {
            queue.data.push(None);
        });

        queue
    }
}

impl<T> Default for ArrayQueue<T> {
    fn default() -> Self {
        Self::new()
    }
}

impl<T> Queue<T> for ArrayQueue<T> {
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

        self.data[self.front]
            .as_ref()
            .expect("element should exist")
    }

    fn enqueue(&mut self, element: T) {
        if self.logical_size == self.physical_size {
            let mut new_capacity = Self::DEFAULT_CAPACITY;

            if self.physical_size >= Self::DEFAULT_CAPACITY {
                new_capacity = self.physical_size + (self.physical_size >> 1)
            }

            let mut temp = Vec::with_capacity(new_capacity);

            let mut cursor = self.front;

            (0..self.logical_size).for_each(|_| {
                if cursor == self.physical_size {
                    cursor = 0;
                }
                temp.push(self.data[cursor].take());
                cursor += 1;
            });
            (self.logical_size..new_capacity).for_each(|_| {
                temp.push(None);
            });

            self.data = temp;
            self.front = 0;
            self.physical_size = new_capacity;
        }

        self.data[(self.front + self.logical_size) % self.physical_size] = Some(element);

        self.logical_size += 1;
    }

    fn dequeue(&mut self) {
        if self.logical_size == 0 {
            panic!("[PANIC - NoSuchElement]");
        }

        self.data[self.front].take();

        self.front = (self.front + 1) % self.physical_size;

        self.logical_size -= 1;
    }
}

impl<T> Resizable for ArrayQueue<T> {
    fn capacity(&self) -> usize {
        self.physical_size
    }

    fn shrink(&mut self) {
        let mut temp = Vec::with_capacity(self.logical_size);

        let mut cursor = self.front;

        (0..self.logical_size).for_each(|_| {
            if cursor == self.physical_size {
                cursor = 0;
            }
            temp.push(self.data[cursor].take());
            cursor += 1;
        });

        self.data = temp;
        self.front = 0;
        self.physical_size = self.logical_size;
    }
}

#[cfg(test)]
pub(crate) mod tests;
