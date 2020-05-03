#ifndef FUN_VAC_DATA_STRUCTURES_DOUBLY_LINKED_LIST_H
#define FUN_VAC_DATA_STRUCTURES_DOUBLY_LINKED_LIST_H 1

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#ifndef DOUBLY_LINKED_LIST_TYPE
typedef int doubly_linked_list_type;
#else
typedef DOUBLY_LINKED_LIST_TYPE doubly_linked_list_type;
#endif

typedef struct DoublyLinkedListNode {
    doubly_linked_list_type data;
    struct DoublyLinkedListNode *next;
    struct DoublyLinkedListNode *prev;
} DoublyLinkedListNode;

typedef struct {
    DoublyLinkedListNode *head;
    DoublyLinkedListNode *tail;
    size_t length;
} DoublyLinkedList;

DoublyLinkedList *DoublyLinkedList_new(void) {
    DoublyLinkedList *list = (DoublyLinkedList *) malloc(sizeof(DoublyLinkedList));

    list->head = NULL;
    list->tail = NULL;
    list->length = 0;

    return list;
}

size_t DoublyLinkedList_size(DoublyLinkedList *list) {
    return list->length;
}

bool DoublyLinkedList_isEmpty(DoublyLinkedList *list) {
    return list->length == 0;
}

doubly_linked_list_type DoublyLinkedList_get(DoublyLinkedList *list, size_t index) {
    if (index < 0 || index >= list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    DoublyLinkedListNode *cursor;

    if (index < list->length >> 1) {
        cursor = list->head;
        for (size_t i = 0; i < index; i++) {
            cursor = cursor->next;
        }
    } else {
        cursor = list->tail;
        for (size_t i = list->length - 1; i > index; i--) {
            cursor = cursor->prev;
        }
    }

    return cursor->data;
}

void DoublyLinkedList_set(DoublyLinkedList *list, size_t index, doubly_linked_list_type element) {
    if (index < 0 || index >= list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    DoublyLinkedListNode *cursor;

    if (index < list->length >> 1) {
        cursor = list->head;
        for (size_t i = 0; i < index; i++) {
            cursor = cursor->next;
        }
    } else {
        cursor = list->tail;
        for (size_t i = list->length - 1; i > index; i--) {
            cursor = cursor->prev;
        }
    }

    cursor->data = element;
}

void DoublyLinkedList_insert(DoublyLinkedList *list, size_t index, doubly_linked_list_type element) {
    if (index < 0 || index > list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    DoublyLinkedListNode *node = (DoublyLinkedListNode *) malloc(sizeof(DoublyLinkedListNode));

    node->data = element;
    node->next = NULL;
    node->prev = NULL;

    if (index == 0) {
        if (list->length != 0) {
            node->next = list->head;
            list->head->prev = node;
        } else {
            list->tail = node;
        }
        list->head = node;
    } else if (index == list->length) {
        node->prev = list->tail;
        list->tail->next = node;
        list->tail = node;
    } else {
        DoublyLinkedListNode *cursor;
        if (index < list->length >> 1) {
            cursor = list->head;
            for (size_t i = 0; i < index; i++) {
                cursor = cursor->next;
            }
        } else {
            cursor = list->tail;
            for (size_t i = list->length - 1; i > index; i--) {
                cursor = cursor->prev;
            }
        }
        node->next = cursor;
        node->prev = cursor->prev;
        cursor->prev->next = node;
        cursor->prev = node;
    }

    list->length += 1;
}

void DoublyLinkedList_remove(DoublyLinkedList *list, size_t index) {
    if (index < 0 || index >= list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    DoublyLinkedListNode *target;

    if (index == 0) {
        target = list->head;
        if (list->length == 1) {
            list->head = NULL;
            list->tail = NULL;
        } else {
            target->next->prev = NULL;
            list->head = target->next;
        }
    } else if (index == list->length - 1) {
        target = list->tail;
        target->prev->next = NULL;
        list->tail = target->prev;
    } else {
        if (index < list->length >> 1) {
            target = list->head;
            for (size_t i = 0; i < index; i++) {
                target = target->next;
            }
        } else {
            target = list->tail;
            for (size_t i = list->length - 1; i > index; i--) {
                target = target->prev;
            }
        }
        target->prev->next = target->next;
        target->next->prev = target->prev;
    }

    target->data = (doubly_linked_list_type) 0;
    free(target);

    list->length -= 1;
}

doubly_linked_list_type DoublyLinkedList_front(DoublyLinkedList *list) {
    return DoublyLinkedList_get(list, 0);
}

doubly_linked_list_type DoublyLinkedList_back(DoublyLinkedList *list) {
    return DoublyLinkedList_get(list, list->length - 1);
}

void DoublyLinkedList_prepend(DoublyLinkedList *list, doubly_linked_list_type element) {
    DoublyLinkedList_insert(list, 0, element);
}

void DoublyLinkedList_append(DoublyLinkedList *list, doubly_linked_list_type element) {
    DoublyLinkedList_insert(list, list->length, element);
}

void DoublyLinkedList_poll(DoublyLinkedList *list) {
    DoublyLinkedList_remove(list, 0);
}

void DoublyLinkedList_eject(DoublyLinkedList *list) {
    DoublyLinkedList_remove(list, list->length - 1);
}

#endif
