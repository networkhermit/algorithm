#ifndef VAC_FUN_DATA_STRUCTURES_CIRCULARLY_LINKED_LIST_H
#define VAC_FUN_DATA_STRUCTURES_CIRCULARLY_LINKED_LIST_H 1

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#ifndef CIRCULARLY_LINKED_LIST_TYPE
typedef int circularly_linked_list_type;
#else
typedef CIRCULARLY_LINKED_LIST_TYPE circularly_linked_list_type;
#endif

typedef struct CircularlyLinkedListNode {
    circularly_linked_list_type data;
    struct CircularlyLinkedListNode *next;
} CircularlyLinkedListNode;

typedef struct {
    CircularlyLinkedListNode *tail;
    size_t length;
} CircularlyLinkedList;

CircularlyLinkedList *CircularlyLinkedList_new(void) {
    CircularlyLinkedList *list = (CircularlyLinkedList *) malloc(sizeof(CircularlyLinkedList));

    list->tail = NULL;
    list->length = 0;

    return list;
}

size_t CircularlyLinkedList_size(CircularlyLinkedList *list) {
    return list->length;
}

bool CircularlyLinkedList_isEmpty(CircularlyLinkedList *list) {
    return list->length == 0;
}

circularly_linked_list_type CircularlyLinkedList_get(CircularlyLinkedList *list, size_t index) {
    if (index < 0 || index >= list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    CircularlyLinkedListNode *cursor = list->tail;

    if (index != list->length - 1) {
        for (size_t i = 0; i <= index; i++) {
            cursor = cursor->next;
        }
    }

    return cursor->data;
}

void CircularlyLinkedList_set(CircularlyLinkedList *list, size_t index, circularly_linked_list_type element) {
    if (index < 0 || index >= list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    CircularlyLinkedListNode *cursor = list->tail;

    if (index != list->length - 1) {
        for (size_t i = 0; i <= index; i++) {
            cursor = cursor->next;
        }
    }

    cursor->data = element;
}

void CircularlyLinkedList_insert(CircularlyLinkedList *list, size_t index, circularly_linked_list_type element) {
    if (index < 0 || index > list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    CircularlyLinkedListNode *node = (CircularlyLinkedListNode *) malloc(sizeof(CircularlyLinkedListNode));

    node->data = element;
    node->next = NULL;

    if (index == 0) {
        if (list->length == 0) {
            node->next = node;
            list->tail = node;
        } else {
            node->next = list->tail->next;
            list->tail->next = node;
        }
    } else if (index == list->length) {
        node->next = list->tail->next;
        list->tail->next = node;
        list->tail = node;
    } else {
        CircularlyLinkedListNode *cursor = list->tail;
        for (size_t i = 0, bound = index - 1; i <= bound; i++) {
            cursor = cursor->next;
        }
        node->next = cursor->next;
        cursor->next = node;
    }

    list->length++;
}

void CircularlyLinkedList_remove(CircularlyLinkedList *list, size_t index) {
    if (index < 0 || index >= list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    CircularlyLinkedListNode *target;

    if (index == 0) {
        target = list->tail->next;
        if (list->length == 1) {
            list->tail = NULL;
        } else {
            list->tail->next = target->next;
        }
    } else {
        CircularlyLinkedListNode *cursor = list->tail;
        for (size_t i = 0, bound = index - 1; i <= bound; i++) {
            cursor = cursor->next;
        }
        target = cursor->next;
        cursor->next = target->next;
        if (index == list->length - 1) {
            list->tail = cursor;
        }
    }

    target->data = (circularly_linked_list_type) 0;
    free(target);

    list->length--;
}

circularly_linked_list_type CircularlyLinkedList_front(CircularlyLinkedList *list) {
    return CircularlyLinkedList_get(list, 0);
}

circularly_linked_list_type CircularlyLinkedList_back(CircularlyLinkedList *list) {
    return CircularlyLinkedList_get(list, list->length - 1);
}

void CircularlyLinkedList_prepend(CircularlyLinkedList *list, circularly_linked_list_type element) {
    CircularlyLinkedList_insert(list, 0, element);
}

void CircularlyLinkedList_append(CircularlyLinkedList *list, circularly_linked_list_type element) {
    CircularlyLinkedList_insert(list, list->length, element);
}

void CircularlyLinkedList_poll(CircularlyLinkedList *list) {
    CircularlyLinkedList_remove(list, 0);
}

void CircularlyLinkedList_eject(CircularlyLinkedList *list) {
    CircularlyLinkedList_remove(list, list->length - 1);
}

#endif
