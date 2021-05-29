#ifndef VAC_FUN_DATA_STRUCTURES_SINGLY_LINKED_LIST_HPP
#define VAC_FUN_DATA_STRUCTURES_SINGLY_LINKED_LIST_HPP 1

#include <stdexcept>

namespace SinglyLinkedList {
    template <typename E>
    class SinglyLinkedList {
    private:
        class Node {
        public:
            E data;
            Node *next = nullptr;

        public:
            Node(E element) {
                data = element;
            }

            ~Node() {}
        };

    private:
        Node *head = nullptr;
        std::size_t length = 0;

    public:
        SinglyLinkedList() {}

        ~SinglyLinkedList() {}

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

            auto cursor = head;

            for (std::size_t i = 0; i < index; i++) {
                cursor = cursor->next;
            }

            return cursor->data;
        }

        void set(std::size_t index, E element) {
            if (index < 0 || index >= length) {
                throw std::runtime_error("[PANIC - IndexOutOfBounds]");
            }

            auto cursor = head;

            for (std::size_t i = 0; i < index; i++) {
                cursor = cursor->next;
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
                }
                head = node;
            } else {
                auto cursor = head;
                for (std::size_t i = 0, bound = index - 1; i < bound; i++) {
                    cursor = cursor->next;
                }
                node->next = cursor->next;
                cursor->next = node;
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
                } else {
                    head = target->next;
                }
            } else {
                auto cursor = head;
                for (std::size_t i = 0, bound = index - 1; i < bound; i++) {
                    cursor = cursor->next;
                }
                target = cursor->next;
                cursor->next = target->next;
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
