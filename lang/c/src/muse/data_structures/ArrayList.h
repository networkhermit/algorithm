#ifndef MUSE_DATA_STRUCTURES_ARRAY_LIST_H
#define MUSE_DATA_STRUCTURES_ARRAY_LIST_H 1

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#ifndef ARRAY_LIST_TYPE
typedef int array_list_type;
#else
typedef ARRAY_LIST_TYPE array_list_type;
#endif

#define ARRAY_LIST_DEFAULT_CAPACITY 64

typedef struct {
    array_list_type *data;
    size_t logicalSize;
    size_t physicalSize;
} ArrayList;

ArrayList *ArrayList_new(size_t physicalSize) {
    ArrayList *list = (ArrayList *) malloc(sizeof(ArrayList));

    list->data = NULL;
    list->logicalSize = 0;
    list->physicalSize = ARRAY_LIST_DEFAULT_CAPACITY;

    if (physicalSize > 1) {
        list->physicalSize = physicalSize;
    }
    list->data = (array_list_type *) malloc(list->physicalSize * sizeof(array_list_type));

    return list;
}

size_t ArrayList_size(ArrayList *list) {
    return list->logicalSize;
}

bool ArrayList_isEmpty(ArrayList *list) {
    return list->logicalSize == 0;
}

array_list_type ArrayList_get(ArrayList *list, size_t index) {
    if (index < 0 || index >= list->logicalSize) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    return list->data[index];
}

void ArrayList_set(ArrayList *list, size_t index, array_list_type element) {
    if (index < 0 || index >= list->logicalSize) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    list->data[index] = element;
}

void ArrayList_insert(ArrayList *list, size_t index, array_list_type element) {
    if (index < 0 || index > list->logicalSize) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    if (list->logicalSize == list->physicalSize) {
        size_t newCapacity = ARRAY_LIST_DEFAULT_CAPACITY;

        if (list->physicalSize >= ARRAY_LIST_DEFAULT_CAPACITY) {
            newCapacity = list->physicalSize + (list->physicalSize >> 1);
        }

        array_list_type *temp = (array_list_type *) malloc(newCapacity * sizeof(array_list_type));

        for (size_t i = 0, length = list->logicalSize; i < length; i++) {
            temp[i] = list->data[i];
        }

        free(list->data);

        list->data = temp;
        list->physicalSize = newCapacity;
    }

    for (size_t i = list->logicalSize; i > index; i--) {
        list->data[i] = list->data[i - 1];
    }

    list->data[index] = element;

    list->logicalSize++;
}

void ArrayList_remove(ArrayList *list, size_t index) {
    if (index < 0 || index >= list->logicalSize) {
        fprintf(stderr, "[PANIC - IndexOutOfBounds]\n");
        exit(EXIT_FAILURE);
    }

    for (size_t i = index + 1; i < list->logicalSize; i++) {
        list->data[i - 1] = list->data[i];
    }

    list->logicalSize--;

    list->data[list->logicalSize] = (array_list_type) 0;
}

array_list_type ArrayList_front(ArrayList *list) {
    return ArrayList_get(list, 0);
}

array_list_type ArrayList_back(ArrayList *list) {
    return ArrayList_get(list, list->logicalSize - 1);
}

void LinkedList_prepend(ArrayList *list, array_list_type element) {
    ArrayList_insert(list, 0, element);
}

void ArrayList_append(ArrayList *list, array_list_type element) {
    ArrayList_insert(list, list->logicalSize, element);
}

void ArrayList_poll(ArrayList *list) {
    ArrayList_remove(list, 0);
}

void ArrayList_eject(ArrayList *list) {
    ArrayList_remove(list, list->logicalSize - 1);
}

size_t ArrayList_capacity(ArrayList *list) {
    return list->physicalSize;
}

void ArrayList_shrink(ArrayList *list) {
    array_list_type *temp = (array_list_type *) malloc(list->logicalSize * sizeof(array_list_type));

    for (size_t i = 0, length = list->logicalSize; i < length; i++) {
        temp[i] = list->data[i];
    }

    free(list->data);

    list->data = temp;
    list->physicalSize = list->logicalSize;
}

#endif
