#ifndef VAC_FUN_DATA_STRUCTURES_LINKED_STACK_HPP
#define VAC_FUN_DATA_STRUCTURES_LINKED_STACK_HPP 1

#include <stdexcept>

namespace LinkedStack {
    template <typename E>
    class LinkedStack {
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
        Node *tail = nullptr;
        std::size_t length = 0;

    public:
        LinkedStack() {}

        ~LinkedStack() {}

    public:
        std::size_t size() {
            return length;
        }

        bool isEmpty() {
            return length == 0;
        }

        E peek() {
            if (length == 0) {
                throw std::runtime_error("[PANIC - NoSuchElement]");
            }

            return tail->data;
        }

        void push(E element) {
            auto node = new Node(element);

            if (length == 0) {
                head = node;
            } else {
                tail->next = node;
            }

            tail = node;

            length++;
        }

        void pop() {
            if (length == 0) {
                throw std::runtime_error("[PANIC - NoSuchElement]");
            }

            auto target = tail;

            if (length == 1) {
                head = nullptr;
                tail = nullptr;
            } else {
                auto cursor = head;
                for (std::size_t i = 0, bound = length - 2; i < bound; i++) {
                    cursor = cursor->next;
                }
                cursor->next = nullptr;
                tail = cursor;
            }

            target->data = static_cast<E>(0);
            delete target;

            length--;
        }
    };
}

#endif
