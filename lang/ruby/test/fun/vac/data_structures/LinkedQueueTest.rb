require "fun/vac/data_structures/LinkedQueue"
require "fun/vac/util/TestRunner"

def testLinkedQueue()
    size = 8192

    queue = LinkedQueue::LinkedQueue.new()

    for i in 1 .. size
        queue.enqueue(i)
    end

    if queue.size() != size
        return false
    end

    for i in 1 .. size
        if queue.peek() != i
            return false
        end
        queue.dequeue()
    end

    return queue.isEmpty()
end

if __FILE__ == $0
    TestRunner.parseTest(testLinkedQueue())
end
