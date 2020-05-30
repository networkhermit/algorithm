require "fun/vac/data_structures/LinkedQueue"
require "fun/vac/util/TestRunner"

def testLinkedQueue()
    size = 8192

    queue = LinkedQueue::LinkedQueue.new()

    (1 .. size).each do |i|
        queue.enqueue(i)
    end

    if queue.size() != size
        return false
    end

    (1 .. size).each do |i|
        if queue.peek() != i
            return false
        end
        queue.dequeue()
    end

    queue.isEmpty()
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testLinkedQueue())
end
