require "fun/vac/data_structures/ArrayQueue"
require "fun/vac/util/TestRunner"

def testArrayQueue()
    size = 8192

    queue = ArrayQueue::ArrayQueue.new(0)

    for i in 1 .. size
        queue.enqueue(i)
    end

    queue.shrink()

    if queue.size() != size
        return false
    end

    if queue.capacity() != size
        return false
    end

    for i in 1 .. size
        if queue.peek() != i
            return false
        end
        queue.dequeue()
    end

    queue.shrink()

    unless queue.isEmpty()
        return false
    end

    if queue.capacity() != 0
        return false
    end

    return true
end

if __FILE__ == $0
    TestRunner.parseTest(testArrayQueue())
end
