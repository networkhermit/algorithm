#ifndef FUN_VAC_DATA_STRUCTURES_ARRAY_LIST_HPP
#define FUN_VAC_DATA_STRUCTURES_ARRAY_LIST_HPP 1

#include <stdexcept>

namespace ArrayList {

    template <typename E> class ArrayList {

    private:
        static const std::size_t DEFAULT_CAPACITY = 64;

        E *data = nullptr;
        std::size_t logicalSize = 0;
        std::size_t physicalSize = DEFAULT_CAPACITY;

    public:
        ArrayList(std::size_t physicalSize = 0) {
            if (physicalSize > 1) {
                this->physicalSize = physicalSize;
            }
            data = new E[this->physicalSize];
        }

        ~ArrayList() {}

    public:
        std::size_t size() {
            return logicalSize;
        }

        bool isEmpty() {
            return logicalSize == 0;
        }

        E get(std::size_t index) {
            if (index < 0 || index >= logicalSize) {
                throw std::runtime_error("[PANIC - IndexOutOfBounds]");
            }

            return data[index];
        }

        void set(std::size_t index, E element) {
            if (index < 0 || index >= logicalSize) {
                throw std::runtime_error("[PANIC - IndexOutOfBounds]");
            }

            data[index] = element;
        }

        void insert(std::size_t index, E element) {
            if (index < 0 || index > logicalSize) {
                throw std::runtime_error("[PANIC - IndexOutOfBounds]");
            }

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

            for (std::size_t i = logicalSize; i > index; i--) {
                data[i] = data[i - 1];
            }

            data[index] = element;

            logicalSize++;
        }

        void remove(std::size_t index) {
            if (index < 0 || index >= logicalSize) {
                throw std::runtime_error("[PANIC - IndexOutOfBounds]");
            }

            for (std::size_t i = index + 1; i < logicalSize; i++) {
                data[i - 1] = data[i];
            }

            logicalSize--;

            data[logicalSize] = static_cast<E>(0);
        }

        E front() {
            return get(0);
        }

        E back() {
            return get(logicalSize - 1);
        }

        void prepend(E element) {
            insert(0, element);
        }

        void append(E element) {
            insert(logicalSize, element);
        }

        void poll() {
            remove(0);
        }

        void eject() {
            remove(logicalSize - 1);
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
