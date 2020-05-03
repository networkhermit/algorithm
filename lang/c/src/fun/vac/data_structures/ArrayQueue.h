#ifndef FUN_VAC_DATA_STRUCTURES_ARRAY_QUEUE_H
#define FUN_VAC_DATA_STRUCTURES_ARRAY_QUEUE_H 1

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#ifndef ARRAY_QUEUE_TYPE
typedef int array_queue_type;
#else
typedef ARRAY_QUEUE_TYPE array_queue_type;
#endif

#define ARRAY_QUEUE_DEFAULT_CAPACITY 64

typedef struct {
    array_queue_type *data;
    size_t front;
    size_t logicalSize;
    size_t physicalSize;
} ArrayQueue;

ArrayQueue * ArrayQueue_new(size_t physicalSize) {
    ArrayQueue *queue = (ArrayQueue *) malloc(sizeof(ArrayQueue));

    queue->data = NULL;
    queue->front = 0;
    queue->logicalSize = 0;
    queue->physicalSize = ARRAY_QUEUE_DEFAULT_CAPACITY;

    if (physicalSize > 1) {
        queue->physicalSize = physicalSize;
    }
    queue->data = (array_queue_type *) malloc(queue->physicalSize * sizeof(array_queue_type));

    return queue;
}

size_t ArrayQueue_size(ArrayQueue *queue) {
    return queue->logicalSize;
}

bool ArrayQueue_isEmpty(ArrayQueue *queue) {
    return queue->logicalSize == 0;
}

array_queue_type ArrayQueue_peek(ArrayQueue *queue) {
    if (queue->logicalSize == 0) {
        fprintf(stderr, "[PANIC - NoSuchElement]\n");
        exit(EXIT_FAILURE);
    }

    return queue->data[queue->front];
}

void ArrayQueue_enqueue(ArrayQueue *queue, array_queue_type element) {
    if (queue->logicalSize == queue->physicalSize) {
        size_t newCapacity = ARRAY_QUEUE_DEFAULT_CAPACITY;

        if (queue->physicalSize >= ARRAY_QUEUE_DEFAULT_CAPACITY) {
            newCapacity = queue->physicalSize + (queue->physicalSize >> 1);
        }

        array_queue_type *temp = (array_queue_type *) malloc(newCapacity * sizeof(array_queue_type));

        size_t cursor = queue->front;

        for (size_t i = 0, length = queue->logicalSize; i < length; i++) {
            if (cursor == queue->physicalSize) {
                cursor = 0;
            }
            temp[i] = queue->data[cursor];
            cursor += 1;
        }

        free(queue->data);

        queue->data = temp;
        queue->front = 0;
        queue->physicalSize = newCapacity;
    }

    queue->data[(queue->front + queue->logicalSize) % queue->physicalSize] = element;

    queue->logicalSize += 1;
}

void ArrayQueue_dequeue(ArrayQueue *queue) {
    if (queue->logicalSize == 0) {
        fprintf(stderr, "[PANIC - NoSuchElement]\n");
        exit(EXIT_FAILURE);
    }

    queue->data[queue->front] = (array_queue_type) 0;

    queue->front = (queue->front + 1) % queue->physicalSize;

    queue->logicalSize -= 1;
}

size_t ArrayQueue_capacity(ArrayQueue *queue) {
    return queue->physicalSize;
}

void ArrayQueue_shrink(ArrayQueue *queue) {
    array_queue_type *temp = (array_queue_type *) malloc(queue->logicalSize * sizeof(array_queue_type));

    size_t cursor = queue->front;

    for (size_t i = 0, length = queue->logicalSize; i < length; i++) {
        if (cursor == queue->physicalSize) {
            cursor = 0;
        }
        temp[i] = queue->data[cursor];
        cursor += 1;
    }

    free(queue->data);

    queue->data = temp;
    queue->front = 0;
    queue->physicalSize = queue->logicalSize;
}

#endif
