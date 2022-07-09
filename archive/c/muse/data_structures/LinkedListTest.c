#include <muse/data_structures/LinkedList.h>
#include <muse/util/TestRunner.h>

bool testLinkedList(void) {
  size_t size = 8192;

  LinkedList *list = LinkedList_new();

  for (size_t i = 1; i <= size; i++) {
    LinkedList_append(list, (int)i);
  }

  if (LinkedList_size(list) != size) {
    return false;
  }

  for (size_t i = 0; i < size; i++) {
    if ((size_t)LinkedList_get(list, i) != i + 1) {
      return false;
    }
  }

  for (size_t i = 0; i < size; i++) {
    LinkedList_set(list, i, size - i);
  }

  for (size_t i = 0; i < size; i++) {
    if ((size_t)LinkedList_get(list, i) != size - i) {
      return false;
    }
  }

  for (size_t i = 0, j = size - 1; i < j; i++, j--) {
    int x = LinkedList_get(list, i);
    int y = LinkedList_get(list, j);

    LinkedList_remove(list, i);
    LinkedList_insert(list, i, y);
    LinkedList_remove(list, j);
    LinkedList_insert(list, j, x);
  }

  for (size_t i = 0; i < size; i++) {
    if ((size_t)LinkedList_get(list, i) != i + 1) {
      return false;
    }
  }

  for (size_t i = size; i >= 1; i--) {
    if ((size_t)LinkedList_back(list) != i) {
      return false;
    }
    LinkedList_eject(list);
  }

  return LinkedList_isEmpty(list);
}

int main(void) { TestRunner_pick(&testLinkedList); }
