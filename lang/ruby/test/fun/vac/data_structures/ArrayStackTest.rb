require "fun/vac/data_structures/ArrayStack"
require "fun/vac/util/TestRunner"

def testArrayStack()
    size = 8192

    stack = ArrayStack::ArrayStack.new(0)

    for i in 1 .. size
        stack.push(i)
    end

    stack.shrink()

    if stack.size() != size
        return false
    end

    if stack.capacity() != size
        return false
    end

    i = size
    while i > 0
        if stack.peek() != i
            return false
        end
        stack.pop()
        i -= 1
    end

    stack.shrink()

    unless stack.isEmpty()
        return false
    end

    return stack.capacity() == 0
end

if __FILE__ == $0
    TestRunner.parseTest(testArrayStack())
end
