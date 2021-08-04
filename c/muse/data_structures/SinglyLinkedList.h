#ifndef MUSE_DATA_STRUCTURES_SINGLY_LINKED_LIST_H
#define MUSE_DATA_STRUCTURES_SINGLY_LINKED_LIST_H 1

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#ifndef SINGLY_LINKED_LIST_TYPE
typedef int singly_linked_list_type;
#else
typedef SINGLY_LINKED_LIST_TYPE singly_linked_list_type;
#endif

typedef struct SinglyLinkedListNode {
    singly_linked_list_type data;
    struct SinglyLinkedListNode *next;
} SinglyLinkedListNode;

typedef struct {
    SinglyLinkedListNode *head;
    size_t length;
} SinglyLinkedList;

SinglyLinkedList *SinglyLinkedList_new(void) {
    SinglyLinkedList *list = (SinglyLinkedList *) malloc(sizeof(SinglyLinkedList));

    list->head = NULL;
    list->length = 0;

    return list;
}

size_t SinglyLinkedList_size(SinglyLinkedList *list) {
    return list->length;
}

bool SinglyLinkedList_isEmpty(SinglyLinkedList *list) {
    return list->length == 0;
}

singly_linked_list_type SinglyLinkedList_get(SinglyLinkedList *list, size_t index) {
    if (index < 0 || index >= list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    SinglyLinkedListNode *cursor = list->head;

    for (size_t i = 0; i < index; i++) {
        cursor = cursor->next;
    }

    return cursor->data;
}

void SinglyLinkedList_set(SinglyLinkedList *list, size_t index, singly_linked_list_type element) {
    if (index < 0 || index >= list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    SinglyLinkedListNode *cursor = list->head;

    for (size_t i = 0; i < index; i++) {
        cursor = cursor->next;
    }

    cursor->data = element;
}

void SinglyLinkedList_insert(SinglyLinkedList *list, size_t index, singly_linked_list_type element) {
    if (index < 0 || index > list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    SinglyLinkedListNode *node = (SinglyLinkedListNode *) malloc(sizeof(SinglyLinkedListNode));

    node->data = element;
    node->next = NULL;

    if (index == 0) {
        if (list->length != 0) {
            node->next = list->head;
        }
        list->head = node;
    } else {
        SinglyLinkedListNode *cursor = list->head;
        for (size_t i = 0, bound = index - 1; i < bound; i++) {
            cursor = cursor->next;
        }
        node->next = cursor->next;
        cursor->next = node;
    }

    list->length++;
}

void SinglyLinkedList_remove(SinglyLinkedList *list, size_t index) {
    if (index < 0 || index >= list->length) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    SinglyLinkedListNode *target;

    if (index == 0) {
        target = list->head;
        if (list->length == 1) {
            list->head = NULL;
        } else {
            list->head = target->next;
        }
    } else {
        SinglyLinkedListNode *cursor = list->head;
        for (size_t i = 0, bound = index - 1; i < bound; i++) {
            cursor = cursor->next;
        }
        target = cursor->next;
        cursor->next = target->next;
    }

    target->data = (singly_linked_list_type) 0;
    free(target);

    list->length--;
}

singly_linked_list_type SinglyLinkedList_front(SinglyLinkedList *list) {
    return SinglyLinkedList_get(list, 0);
}

singly_linked_list_type SinglyLinkedList_back(SinglyLinkedList *list) {
    return SinglyLinkedList_get(list, list->length - 1);
}

void SinglyLinkedList_prepend(SinglyLinkedList *list, singly_linked_list_type element) {
    SinglyLinkedList_insert(list, 0, element);
}

void SinglyLinkedList_append(SinglyLinkedList *list, singly_linked_list_type element) {
    SinglyLinkedList_insert(list, list->length, element);
}

void SinglyLinkedList_poll(SinglyLinkedList *list) {
    SinglyLinkedList_remove(list, 0);
}

void SinglyLinkedList_eject(SinglyLinkedList *list) {
    SinglyLinkedList_remove(list, list->length - 1);
}

#endif
