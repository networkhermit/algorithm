require "muse/data_structures/ArrayQueue"
require "muse/util/TestRunner"

def testArrayQueue()
    size = 8192

    queue = ArrayQueue::ArrayQueue.new(0)

    1.upto(size) do |i|
        queue.enqueue(i)
    end

    queue.shrink()

    if queue.size() != size
        return false
    end

    if queue.capacity() != size
        return false
    end

    1.upto(size) do |i|
        if queue.peek() != i
            return false
        end
        queue.dequeue()
    end

    queue.shrink()

    unless queue.isEmpty()
        return false
    end

    queue.capacity().zero?
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testArrayQueue())
end
