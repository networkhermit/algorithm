#include <muse/data_structures/LinkedQueue.hpp>
#include <muse/util/TestRunner.hpp>

using namespace std;

bool testLinkedQueue() {
  size_t size = 8192;

  auto queue = new LinkedQueue::LinkedQueue<int>();

  for (size_t i = 1; i <= size; i++) {
    queue->enqueue(static_cast<int>(i));
  }

  if (queue->size() != size) {
    return false;
  }

  for (size_t i = 1; i <= size; i++) {
    if (static_cast<size_t>(queue->peek()) != i) {
      return false;
    }
    queue->dequeue();
  }

  return queue->isEmpty();
}

int main() { TestRunner::pick(&testLinkedQueue); }
