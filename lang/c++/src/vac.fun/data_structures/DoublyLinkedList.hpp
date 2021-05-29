#ifndef VAC_FUN_DATA_STRUCTURES_DOUBLY_LINKED_LIST_HPP
#define VAC_FUN_DATA_STRUCTURES_DOUBLY_LINKED_LIST_HPP 1

#include <stdexcept>

namespace DoublyLinkedList {
    template <typename E>
    class DoublyLinkedList {
    private:
        class Node {
        public:
            E data;
            Node *next = nullptr;
            Node *prev = nullptr;

        public:
            Node(E element) {
                data = element;
            }

            ~Node() {}
        };

    private:
        Node *head = nullptr;
        Node *tail = nullptr;
        std::size_t length = 0;

    public:
        DoublyLinkedList() {}

        ~DoublyLinkedList() {}

    public:
        std::size_t size() {
            return length;
        }

        bool isEmpty() {
            return length == 0;
        }

        E get(std::size_t index) {
            if (index < 0 || index >= length) {
                throw std::runtime_error("[PANIC - IndexOutOfBounds]");
            }

            Node *cursor;

            if (index < length >> 1) {
                cursor = head;
                for (std::size_t i = 0; i < index; i++) {
                    cursor = cursor->next;
                }
            } else {
                cursor = tail;
                for (std::size_t i = length - 1; i > index; i--) {
                    cursor = cursor->prev;
                }
            }

            return cursor->data;
        }

        void set(std::size_t index, E element) {
            if (index < 0 || index >= length) {
                throw std::runtime_error("[PANIC - IndexOutOfBounds]");
            }

            Node *cursor;

            if (index < length >> 1) {
                cursor = head;
                for (std::size_t i = 0; i < index; i++) {
                    cursor = cursor->next;
                }
            } else {
                cursor = tail;
                for (std::size_t i = length - 1; i > index; i--) {
                    cursor = cursor->prev;
                }
            }

            cursor->data = element;
        }

        void insert(std::size_t index, E element) {
            if (index < 0 || index > length) {
                throw std::runtime_error("[PANIC - IndexOutOfBounds]");
            }

            auto node = new Node(element);

            if (index == 0) {
                if (length != 0) {
                    node->next = head;
                    head->prev = node;
                } else {
                    tail = node;
                }
                head = node;
            } else if (index == length) {
                node->prev = tail;
                tail->next = node;
                tail = node;
            } else {
                Node *cursor;
                if (index < length >> 1) {
                    cursor = head;
                    for (std::size_t i = 0; i < index; i++) {
                        cursor = cursor->next;
                    }
                } else {
                    cursor = tail;
                    for (std::size_t i = length - 1; i > index; i--) {
                        cursor = cursor->prev;
                    }
                }
                node->next = cursor;
                node->prev = cursor->prev;
                cursor->prev->next = node;
                cursor->prev = node;
            }

            length++;
        }

        void remove(std::size_t index) {
            if (index < 0 || index >= length) {
                throw std::runtime_error("[PANIC - IndexOutOfBounds]");
            }

            Node *target;

            if (index == 0) {
                target = head;
                if (length == 1) {
                    head = nullptr;
                    tail = nullptr;
                } else {
                    target->next->prev = nullptr;
                    head = target->next;
                }
            } else if (index == length - 1) {
                target = tail;
                target->prev->next = nullptr;
                tail = target->prev;
            } else {
                if (index < length >> 1) {
                    target = head;
                    for (std::size_t i = 0; i < index; i++) {
                        target = target->next;
                    }
                } else {
                    target = tail;
                    for (std::size_t i = length - 1; i > index; i--) {
                        target = target->prev;
                    }
                }
                target->prev->next = target->next;
                target->next->prev = target->prev;
            }

            target->data = static_cast<E>(0);
            delete target;

            length--;
        }

        E front() {
            return get(0);
        }

        E back() {
            return get(length - 1);
        }

        void prepend(E element) {
            insert(0, element);
        }

        void append(E element) {
            insert(length, element);
        }

        void poll() {
            remove(0);
        }

        void eject() {
            remove(length - 1);
        }
    };
}

#endif
