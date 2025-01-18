#ifndef MUSE_DATA_STRUCTURES_ARRAY_STACK_H
#define MUSE_DATA_STRUCTURES_ARRAY_STACK_H

#include <stdio.h>
#include <stdlib.h>

#ifndef ARRAY_STACK_TYPE
typedef int array_stack_type;
#else
typedef ARRAY_STACK_TYPE array_stack_type;
#endif

#define ARRAY_STACK_DEFAULT_CAPACITY 64

typedef struct {
  array_stack_type *data;
  size_t logicalSize;
  size_t physicalSize;
} ArrayStack;

ArrayStack *ArrayStack_new(size_t physicalSize) {
  ArrayStack *stack = (ArrayStack *)malloc(sizeof(ArrayStack));

  stack->data = nullptr;
  stack->logicalSize = 0;
  stack->physicalSize = ARRAY_STACK_DEFAULT_CAPACITY;

  if (physicalSize > 1) {
    stack->physicalSize = physicalSize;
  }
  stack->data = (array_stack_type *)malloc(stack->physicalSize *
                                           sizeof(array_stack_type));

  return stack;
}

size_t ArrayStack_size(ArrayStack *stack) { return stack->logicalSize; }

bool ArrayStack_isEmpty(ArrayStack *stack) { return stack->logicalSize == 0; }

array_stack_type ArrayStack_peek(ArrayStack *stack) {
  if (stack->logicalSize == 0) {
    fprintf(stderr, "[PANIC - NoSuchElement]\n");
    exit(EXIT_FAILURE);
  }

  return stack->data[stack->logicalSize - 1];
}

void ArrayStack_push(ArrayStack *stack, array_stack_type element) {
  if (stack->logicalSize == stack->physicalSize) {
    size_t newCapacity = ARRAY_STACK_DEFAULT_CAPACITY;

    if (stack->physicalSize >= ARRAY_STACK_DEFAULT_CAPACITY) {
      newCapacity = stack->physicalSize + (stack->physicalSize >> 1);
    }

    array_stack_type *temp =
        (array_stack_type *)malloc(newCapacity * sizeof(array_stack_type));

    for (size_t i = 0, length = stack->logicalSize; i < length; i++) {
      temp[i] = stack->data[i];
    }

    free(stack->data);

    stack->data = temp;
    stack->physicalSize = newCapacity;
  }

  stack->data[stack->logicalSize] = element;

  stack->logicalSize++;
}

void ArrayStack_pop(ArrayStack *stack) {
  if (stack->logicalSize == 0) {
    fprintf(stderr, "[PANIC - NoSuchElement]\n");
    exit(EXIT_FAILURE);
  }

  stack->logicalSize--;

  stack->data[stack->logicalSize] = (array_stack_type)0;
}

size_t ArrayStack_capacity(ArrayStack *stack) { return stack->physicalSize; }

void ArrayStack_shrink(ArrayStack *stack) {
  array_stack_type *temp =
      (array_stack_type *)malloc(stack->logicalSize * sizeof(array_stack_type));

  for (size_t i = 0, length = stack->logicalSize; i < length; i++) {
    temp[i] = stack->data[i];
  }

  free(stack->data);

  stack->data = temp;
  stack->physicalSize = stack->logicalSize;
}

#endif
