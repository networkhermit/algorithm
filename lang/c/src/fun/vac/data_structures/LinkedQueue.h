#ifndef FUN_VAC_DATA_STRUCTURES_LINKED_QUEUE_H
#define FUN_VAC_DATA_STRUCTURES_LINKED_QUEUE_H 1

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#ifndef LINKED_QUEUE_TYPE
typedef int linked_queue_type;
#else
typedef LINKED_QUEUE_TYPE linked_queue_type;
#endif

typedef struct LinkedQueueNode {
    linked_queue_type data;
    struct LinkedQueueNode *next;
} LinkedQueueNode;

typedef struct {
    LinkedQueueNode *head;
    LinkedQueueNode *tail;
    size_t length;
} LinkedQueue;

LinkedQueue *LinkedQueue_new(void) {
    LinkedQueue *queue = (LinkedQueue *) malloc(sizeof(LinkedQueue));

    queue->head = NULL;
    queue->tail = NULL;
    queue->length = 0;

    return queue;
}

size_t LinkedQueue_size(LinkedQueue *queue) {
    return queue->length;
}

bool LinkedQueue_isEmpty(LinkedQueue *queue) {
    return queue->length == 0;
}

linked_queue_type LinkedQueue_peek(LinkedQueue *queue) {
    if (queue->length == 0) {
        fprintf(stderr, "[PANIC - NoSuchElement]\n");
        exit(EXIT_FAILURE);
    }

    return queue->head->data;
}

void LinkedQueue_enqueue(LinkedQueue *queue, linked_queue_type element) {
    LinkedQueueNode *node = (LinkedQueueNode *) malloc(sizeof(LinkedQueueNode));

    node->data = element;
    node->next = NULL;

    if (queue->length == 0) {
        queue->head = node;
    } else {
        queue->tail->next = node;
    }

    queue->tail = node;

    queue->length++;
}

void LinkedQueue_dequeue(LinkedQueue *queue) {
    if (queue->length == 0) {
        fprintf(stderr, "[PANIC - NoSuchElement]\n");
        exit(EXIT_FAILURE);
    }

    LinkedQueueNode *target = queue->head;

    if (queue->length == 1) {
        queue->head = NULL;
        queue->tail = NULL;
    } else {
        queue->head = target->next;
    }

    target->data = (linked_queue_type) 0;
    free(target);

    queue->length--;
}

#endif
