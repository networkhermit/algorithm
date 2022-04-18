require "muse/data_structures/LinkedQueue"
require "muse/util/TestRunner"

def testLinkedQueue
  size = 8192

  queue = LinkedQueue::LinkedQueue.new

  1.upto(size) do |i|
    queue.enqueue(i)
  end

  return false if queue.size != size

  1.upto(size) do |i|
    return false if queue.peek != i

    queue.dequeue
  end

  queue.isEmpty
end

TestRunner.pick(testLinkedQueue) if __FILE__ == $PROGRAM_NAME
