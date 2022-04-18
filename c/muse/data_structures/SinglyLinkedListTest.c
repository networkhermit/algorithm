#include <muse/data_structures/SinglyLinkedList.h>
#include <muse/util/TestRunner.h>

bool testSinglyLinkedList(void) {
  size_t size = 8192;

  SinglyLinkedList *list = SinglyLinkedList_new();

  for (size_t i = 1; i <= size; i++) {
    SinglyLinkedList_append(list, (int)i);
  }

  if (SinglyLinkedList_size(list) != size) {
    return false;
  }

  for (size_t i = 0; i < size; i++) {
    if ((size_t)SinglyLinkedList_get(list, i) != i + 1) {
      return false;
    }
  }

  for (size_t i = 0; i < size; i++) {
    SinglyLinkedList_set(list, i, size - i);
  }

  for (size_t i = 0; i < size; i++) {
    if ((size_t)SinglyLinkedList_get(list, i) != size - i) {
      return false;
    }
  }

  for (size_t i = 0, j = size - 1; i < j; i++, j--) {
    int x = SinglyLinkedList_get(list, i);
    int y = SinglyLinkedList_get(list, j);

    SinglyLinkedList_remove(list, i);
    SinglyLinkedList_insert(list, i, y);
    SinglyLinkedList_remove(list, j);
    SinglyLinkedList_insert(list, j, x);
  }

  for (size_t i = 0; i < size; i++) {
    if ((size_t)SinglyLinkedList_get(list, i) != i + 1) {
      return false;
    }
  }

  for (size_t i = size; i >= 1; i--) {
    if ((size_t)SinglyLinkedList_back(list) != i) {
      return false;
    }
    SinglyLinkedList_eject(list);
  }

  return SinglyLinkedList_isEmpty(list);
}

int main(void) { TestRunner_pick(&testSinglyLinkedList); }
