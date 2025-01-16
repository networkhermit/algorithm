use std::ptr::NonNull;

use super::List;

#[derive(Debug)]
pub struct Node<T> {
    data: Option<T>,
    next: Option<NonNull<Node<T>>>,
}

#[derive(Debug)]
pub struct LinkedList<T> {
    head: Option<NonNull<Node<T>>>,
    tail: Option<NonNull<Node<T>>>,
    length: usize,
}

impl<T> LinkedList<T> {
    pub fn new() -> Self {
        Self {
            head: None,
            tail: None,
            length: 0,
        }
    }
}

impl<T> Default for LinkedList<T> {
    fn default() -> Self {
        Self::new()
    }
}

impl<T> Drop for LinkedList<T> {
    fn drop(&mut self) {
        while !self.is_empty() {
            self.eject();
        }
    }
}

impl<T> List<T> for LinkedList<T> {
    fn size(&self) -> usize {
        self.length
    }

    fn is_empty(&self) -> bool {
        self.length == 0
    }

    fn get(&self, index: usize) -> &T {
        if index >= self.length {
            panic!("[PANIC - IndexOutOfBounds]");
        }

        let mut cursor;

        if index == self.length - 1 {
            cursor = &self.tail;
        } else {
            cursor = &self.head;
            (0..index).for_each(|_| {
                cursor = unsafe { &cursor.as_ref().expect("element should exist").as_ref().next };
            });
        }

        unsafe {
            cursor
                .as_ref()
                .expect("element should exist")
                .as_ref()
                .data
                .as_ref()
                .expect("element should exist")
        }
    }

    fn set(&mut self, index: usize, element: T) {
        if index >= self.length {
            panic!("[PANIC - IndexOutOfBounds]");
        }

        let mut cursor;

        if index == self.length - 1 {
            cursor = &mut self.tail;
        } else {
            cursor = &mut self.head;
            for _ in 0..index {
                cursor =
                    unsafe { &mut cursor.as_mut().expect("element should exist").as_mut().next };
            }
        }

        unsafe {
            cursor.as_mut().expect("element should exist").as_mut().data = Some(element);
        }
    }

    fn insert(&mut self, index: usize, element: T) {
        if index > self.length {
            panic!("[PANIC - IndexOutOfBounds]");
        }

        let node = Box::new(Node {
            data: Some(element),
            next: None,
        });

        let mut ptr: NonNull<Node<T>> = Box::leak(node).into();

        if index == 0 {
            if self.length != 0 {
                unsafe {
                    ptr.as_mut().next = self.head;
                }
            } else {
                self.tail = Some(ptr);
            }
            self.head = Some(ptr);
        } else if index == self.length {
            unsafe {
                self.tail
                    .as_mut()
                    .expect("element should exist")
                    .as_mut()
                    .next = Some(ptr);
            }
            self.tail = Some(ptr);
        } else {
            let mut cursor = &mut self.head;
            for _ in 0..index - 1 {
                unsafe {
                    cursor = &mut cursor.as_mut().expect("element should exist").as_mut().next;
                }
            }
            unsafe {
                ptr.as_mut().next = cursor.as_mut().expect("element should exist").as_mut().next;
                cursor.as_mut().expect("element should exist").as_mut().next = Some(ptr);
            }
        }

        self.length += 1;
    }

    fn remove(&mut self, index: usize) {
        if index >= self.length {
            panic!("[PANIC - IndexOutOfBounds]");
        }

        if index == 0 {
            let mut target = unsafe {
                Box::from_raw(self.head.as_mut().expect("element should exist").as_mut())
            };
            target.as_mut().data.take();
            if self.length == 1 {
                self.head.take();
                self.tail.take();
            } else {
                unsafe {
                    self.head = self
                        .head
                        .as_mut()
                        .expect("element should exist")
                        .as_mut()
                        .next;
                }
            }
        } else {
            let mut cursor = &mut self.head;
            for _ in 0..index - 1 {
                cursor =
                    unsafe { &mut cursor.as_mut().expect("element should exist").as_mut().next };
            }
            let mut target = unsafe {
                let ptr = &mut cursor.as_mut().expect("element should exist").as_mut().next;
                Box::from_raw(ptr.as_mut().expect("element should exist").as_mut())
            };
            target.as_mut().data.take();
            unsafe {
                cursor.as_mut().expect("element should exist").as_mut().next = target.as_mut().next;
            }
            if index == self.length - 1 {
                self.tail = *cursor;
            }
        }

        self.length -= 1;
    }

    fn front(&self) -> &T {
        self.get(0)
    }

    fn back(&self) -> &T {
        self.get(self.length - 1)
    }

    fn prepend(&mut self, element: T) {
        self.insert(0, element)
    }

    fn append(&mut self, element: T) {
        self.insert(self.length, element)
    }

    fn poll(&mut self) {
        self.remove(0)
    }

    fn eject(&mut self) {
        self.remove(self.length - 1)
    }
}

#[cfg(test)]
pub(crate) mod tests;
