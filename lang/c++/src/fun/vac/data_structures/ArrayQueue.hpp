#ifndef FUN_VAC_DATA_STRUCTURES_ARRAY_QUEUE_HPP
#define FUN_VAC_DATA_STRUCTURES_ARRAY_QUEUE_HPP 1

#include <stdexcept>

namespace ArrayQueue {

    template <typename E> class ArrayQueue {

    private:
        static const std::size_t DEFAULT_CAPACITY = 64;

        E *data = nullptr;
        std::size_t front = 0;
        std::size_t logicalSize = 0;
        std::size_t physicalSize = DEFAULT_CAPACITY;

    public:
        ArrayQueue(std::size_t physicalSize = 0) {
            if (physicalSize > 1) {
                this->physicalSize = physicalSize;
            }
            data = new E[this->physicalSize];
        }

        ~ArrayQueue() {}

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

            return data[front];
        }

        void enqueue(E element) {
            if (logicalSize == physicalSize) {
                std::size_t newCapacity = DEFAULT_CAPACITY;

                if (physicalSize >= DEFAULT_CAPACITY) {
                    newCapacity = physicalSize + (physicalSize >> 1);
                }

                E *temp = new E[newCapacity];

                std::size_t cursor = front;

                for (std::size_t i = 0, length = logicalSize; i < length; i++) {
                    if (cursor == physicalSize) {
                        cursor = 0;
                    }
                    temp[i] = data[cursor];
                    cursor += 1;
                }

                delete[] data;

                data = temp;
                front = 0;
                physicalSize = newCapacity;
            }

            data[(front + logicalSize) % physicalSize] = element;

            logicalSize += 1;
        }

        void dequeue() {
            if (logicalSize == 0) {
                throw std::runtime_error("[PANIC - NoSuchElement]");
            }

            data[front] = static_cast<E>(0);

            front = (front + 1) % physicalSize;

            logicalSize -= 1;
        }

        std::size_t capacity() {
            return physicalSize;
        }

        void shrink() {
            E *temp = new E[logicalSize];

            std::size_t cursor = front;

            for (std::size_t i = 0, length = logicalSize; i < length; i++) {
                if (cursor == physicalSize) {
                    cursor = 0;
                }
                temp[i] = data[cursor];
                cursor += 1;
            }

            delete[] data;

            data = temp;
            front = 0;
            physicalSize = logicalSize;
        }
    };
}

#endif
