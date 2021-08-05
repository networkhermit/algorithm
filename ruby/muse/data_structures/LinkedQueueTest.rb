require "muse/data_structures/LinkedQueue"
require "muse/util/TestRunner"

def testLinkedQueue()
    size = 8192

    queue = LinkedQueue::LinkedQueue.new()

    1.upto(size) do |i|
        queue.enqueue(i)
    end

    if queue.size() != size
        return false
    end

    1.upto(size) do |i|
        if queue.peek() != i
            return false
        end
        queue.dequeue()
    end

    queue.isEmpty()
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.pick(testLinkedQueue())
end
