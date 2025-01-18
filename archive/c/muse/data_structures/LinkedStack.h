#ifndef MUSE_DATA_STRUCTURES_LINKED_STACK_H
#define MUSE_DATA_STRUCTURES_LINKED_STACK_H

#include <stdio.h>
#include <stdlib.h>

#ifndef LINKED_STACK_TYPE
typedef int linked_stack_type;
#else
typedef LINKED_STACK_TYPE linked_stack_type;
#endif

typedef struct LinkedStackNode {
  linked_stack_type data;
  struct LinkedStackNode *next;
} LinkedStackNode;

typedef struct {
  LinkedStackNode *head;
  LinkedStackNode *tail;
  size_t length;
} LinkedStack;

LinkedStack *LinkedStack_new(void) {
  LinkedStack *stack = (LinkedStack *)malloc(sizeof(LinkedStack));

  stack->head = nullptr;
  stack->tail = nullptr;
  stack->length = 0;

  return stack;
}

size_t LinkedStack_size(LinkedStack *stack) { return stack->length; }

bool LinkedStack_isEmpty(LinkedStack *stack) { return stack->length == 0; }

linked_stack_type LinkedStack_peek(LinkedStack *stack) {
  if (stack->length == 0) {
    fprintf(stderr, "[PANIC - NoSuchElement]\n");
    exit(EXIT_FAILURE);
  }

  return stack->tail->data;
}

void LinkedStack_push(LinkedStack *stack, linked_stack_type element) {
  LinkedStackNode *node = (LinkedStackNode *)malloc(sizeof(LinkedStackNode));

  node->data = element;
  node->next = nullptr;

  if (stack->length == 0) {
    stack->head = node;
  } else {
    stack->tail->next = node;
  }

  stack->tail = node;

  stack->length++;
}

void LinkedStack_pop(LinkedStack *stack) {
  if (stack->length == 0) {
    fprintf(stderr, "[PANIC - NoSuchElement]\n");
    exit(EXIT_FAILURE);
  }

  LinkedStackNode *target = stack->tail;

  if (stack->length == 1) {
    stack->head = nullptr;
    stack->tail = nullptr;
  } else {
    LinkedStackNode *cursor = stack->head;
    for (size_t i = 0, bound = stack->length - 2; i < bound; i++) {
      cursor = cursor->next;
    }
    cursor->next = nullptr;
    stack->tail = cursor;
  }

  target->data = (linked_stack_type)0;
  free(target);

  stack->length--;
}

#endif
