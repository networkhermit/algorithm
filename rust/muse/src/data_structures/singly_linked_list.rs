use super::List;

#[derive(Debug)]
pub struct Node<T> {
    data: Option<T>,
    next: Option<Box<Node<T>>>,
}

#[derive(Debug)]
pub struct SinglyLinkedList<T> {
    head: Option<Box<Node<T>>>,
    length: usize,
}

impl<T> SinglyLinkedList<T> {
    pub fn new() -> Self {
        Self {
            head: None,
            length: 0,
        }
    }
}

impl<T> Default for SinglyLinkedList<T> {
    fn default() -> Self {
        Self::new()
    }
}

impl<T> List<T> for SinglyLinkedList<T> {
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

        let mut cursor = &self.head;

        (0..index).for_each(|_| {
            cursor = &cursor.as_ref().expect("element should exist").next;
        });

        cursor
            .as_ref()
            .expect("element should exist")
            .data
            .as_ref()
            .expect("element should exist")
    }

    fn set(&mut self, index: usize, element: T) {
        if index >= self.length {
            panic!("[PANIC - IndexOutOfBounds]");
        }

        let mut cursor = &mut self.head;

        for _ in 0..index {
            cursor = &mut cursor.as_mut().expect("element should exist").next;
        }

        cursor.as_mut().expect("element should exist").data = Some(element);
    }

    fn insert(&mut self, index: usize, element: T) {
        if index > self.length {
            panic!("[PANIC - IndexOutOfBounds]");
        }

        let mut node = Box::new(Node {
            data: Some(element),
            next: None,
        });

        if index == 0 {
            if self.length != 0 {
                node.next = self.head.take();
            }
            self.head = Some(node);
        } else {
            let mut cursor = &mut self.head;
            for _ in 0..index - 1 {
                cursor = &mut cursor.as_mut().expect("element should exist").next;
            }
            node.next = cursor.as_mut().expect("element should exist").next.take();
            cursor.as_mut().expect("element should exist").next = Some(node);
        }

        self.length += 1;
    }

    fn remove(&mut self, index: usize) {
        if index >= self.length {
            panic!("[PANIC - IndexOutOfBounds]");
        }

        if index == 0 {
            self.head
                .as_mut()
                .expect("element should exist")
                .data
                .take();
            if self.length == 1 {
                self.head.take();
            } else {
                self.head = self
                    .head
                    .as_mut()
                    .expect("element should exist")
                    .next
                    .take();
            }
        } else {
            let mut cursor = &mut self.head;
            for _ in 0..index - 1 {
                cursor = &mut cursor.as_mut().expect("element should exist").next;
            }
            let target = &mut cursor.as_mut().expect("element should exist").next;
            target.as_mut().expect("element should exist").data.take();
            cursor.as_mut().expect("element should exist").next =
                target.as_mut().expect("element should exist").next.take();
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
