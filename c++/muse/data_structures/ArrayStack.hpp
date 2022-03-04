#ifndef MUSE_DATA_STRUCTURES_ARRAY_STACK_HPP
#define MUSE_DATA_STRUCTURES_ARRAY_STACK_HPP

#include <stdexcept>

namespace ArrayStack {
    template <typename E>
    class ArrayStack {
    private:
        static const std::size_t DEFAULT_CAPACITY = 64;

        E *data = nullptr;
        std::size_t logicalSize = 0;
        std::size_t physicalSize = DEFAULT_CAPACITY;

    public:
        ArrayStack(std::size_t physicalSize = 0) {
            if (physicalSize > 1) {
                this->physicalSize = physicalSize;
            }
            data = new E[this->physicalSize];
        }

        ~ArrayStack() {}

    public:
        std::size_t size() {
            return logicalSize;
        }

        bool isEmpty() {
            return logicalSize == 0;
        }

        E peek() {
            if (logicalSize == 0) {
                throw std::runtime_error("[PANIC - NoSuchElement]");
            }

            return data[logicalSize - 1];
        }

        void push(E element) {
            if (logicalSize == physicalSize) {
                std::size_t newCapacity = DEFAULT_CAPACITY;

                if (physicalSize >= DEFAULT_CAPACITY) {
                    newCapacity = physicalSize + (physicalSize >> 1);
                }

                E *temp = new E[newCapacity];

                for (std::size_t i = 0, length = logicalSize; i < length; i++) {
                    temp[i] = data[i];
                }

                delete[] data;

                data = temp;
                physicalSize = newCapacity;
            }

            data[logicalSize] = element;

            logicalSize++;
        }

        void pop() {
            if (logicalSize == 0) {
                throw std::runtime_error("[PANIC - NoSuchElement]");
            }

            logicalSize--;

            data[logicalSize] = static_cast<E>(0);
        }

        std::size_t capacity() {
            return physicalSize;
        }

        void shrink() {
            E *temp = new E[logicalSize];

            for (std::size_t i = 0, length = logicalSize; i < length; i++) {
                temp[i] = data[i];
            }

            delete[] data;

            data = temp;
            physicalSize = logicalSize;
        }
    };
}

#endif
