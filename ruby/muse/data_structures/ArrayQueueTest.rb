require "muse/data_structures/ArrayQueue"
require "muse/util/TestRunner"

def testArrayQueue
  size = 8192

  queue = ArrayQueue::ArrayQueue.new(0)

  1.upto(size) do |i|
    queue.enqueue(i)
  end

  queue.shrink

  return false if queue.size != size

  return false if queue.capacity != size

  1.upto(size) do |i|
    return false if queue.peek != i

    queue.dequeue
  end

  queue.shrink

  return false unless queue.isEmpty

  queue.capacity.zero?
end

TestRunner.pick(testArrayQueue) if __FILE__ == $PROGRAM_NAME
