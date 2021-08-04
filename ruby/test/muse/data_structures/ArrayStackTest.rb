require "muse/data_structures/ArrayStack"
require "muse/util/TestRunner"

def testArrayStack()
    size = 8192

    stack = ArrayStack::ArrayStack.new(0)

    1.upto(size) do |i|
        stack.push(i)
    end

    stack.shrink()

    if stack.size() != size
        return false
    end

    if stack.capacity() != size
        return false
    end

    size.downto(1) do |i|
        if stack.peek() != i
            return false
        end
        stack.pop()
    end

    stack.shrink()

    unless stack.isEmpty()
        return false
    end

    stack.capacity().zero?
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testArrayStack())
end
