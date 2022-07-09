#ifndef MUSE_DATA_STRUCTURES_LINKED_LIST_H
#define MUSE_DATA_STRUCTURES_LINKED_LIST_H

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#ifndef LINKED_LIST_TYPE
typedef int linked_list_type;
#else
typedef LINKED_LIST_TYPE linked_list_type;
#endif

typedef struct LinkedListNode {
  linked_list_type data;
  struct LinkedListNode *next;
} LinkedListNode;

typedef struct {
  LinkedListNode *head;
  LinkedListNode *tail;
  size_t length;
} LinkedList;

LinkedList *LinkedList_new(void) {
  LinkedList *list = (LinkedList *)malloc(sizeof(LinkedList));

  list->head = NULL;
  list->tail = NULL;
  list->length = 0;

  return list;
}

size_t LinkedList_size(LinkedList *list) { return list->length; }

bool LinkedList_isEmpty(LinkedList *list) { return list->length == 0; }

linked_list_type LinkedList_get(LinkedList *list, size_t index) {
  if (index < 0 || index >= list->length) {
    fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
    exit(EXIT_FAILURE);
  }

  LinkedListNode *cursor;

  if (index == list->length - 1) {
    cursor = list->tail;
  } else {
    cursor = list->head;
    for (size_t i = 0; i < index; i++) {
      cursor = cursor->next;
    }
  }

  return cursor->data;
}

void LinkedList_set(LinkedList *list, size_t index, linked_list_type element) {
  if (index < 0 || index >= list->length) {
    fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
    exit(EXIT_FAILURE);
  }

  LinkedListNode *cursor;

  if (index == list->length - 1) {
    cursor = list->tail;
  } else {
    cursor = list->head;
    for (size_t i = 0; i < index; i++) {
      cursor = cursor->next;
    }
  }

  cursor->data = element;
}

void LinkedList_insert(LinkedList *list, size_t index,
                       linked_list_type element) {
  if (index < 0 || index > list->length) {
    fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
    exit(EXIT_FAILURE);
  }

  LinkedListNode *node = (LinkedListNode *)malloc(sizeof(LinkedListNode));

  node->data = element;
  node->next = NULL;

  if (index == 0) {
    if (list->length != 0) {
      node->next = list->head;
    } else {
      list->tail = node;
    }
    list->head = node;
  } else if (index == list->length) {
    list->tail->next = node;
    list->tail = node;
  } else {
    LinkedListNode *cursor = list->head;
    for (size_t i = 0, bound = index - 1; i < bound; i++) {
      cursor = cursor->next;
    }
    node->next = cursor->next;
    cursor->next = node;
  }

  list->length++;
}

void LinkedList_remove(LinkedList *list, size_t index) {
  if (index < 0 || index >= list->length) {
    fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
    exit(EXIT_FAILURE);
  }

  LinkedListNode *target;

  if (index == 0) {
    target = list->head;
    if (list->length == 1) {
      list->head = NULL;
      list->tail = NULL;
    } else {
      list->head = target->next;
    }
  } else {
    LinkedListNode *cursor = list->head;
    for (size_t i = 0, bound = index - 1; i < bound; i++) {
      cursor = cursor->next;
    }
    target = cursor->next;
    cursor->next = target->next;
    if (index == list->length - 1) {
      list->tail = cursor;
    }
  }

  target->data = (linked_list_type)0;
  free(target);

  list->length--;
}

linked_list_type LinkedList_front(LinkedList *list) {
  return LinkedList_get(list, 0);
}

linked_list_type LinkedList_back(LinkedList *list) {
  return LinkedList_get(list, list->length - 1);
}

void LinkedList_prepend(LinkedList *list, linked_list_type element) {
  LinkedList_insert(list, 0, element);
}

void LinkedList_append(LinkedList *list, linked_list_type element) {
  LinkedList_insert(list, list->length, element);
}

void LinkedList_poll(LinkedList *list) { LinkedList_remove(list, 0); }

void LinkedList_eject(LinkedList *list) {
  LinkedList_remove(list, list->length - 1);
}

#endif
