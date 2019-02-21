require "fun/vac/data_structures/LinkedStack"
require "fun/vac/util/TestRunner"

def testLinkedStack()
    size = 8192

    stack = LinkedStack::LinkedStack.new()

    for i in 1 .. size
        stack.push(i)
    end

    if stack.size() != size
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

    unless stack.isEmpty()
        return false
    end

    return true
end

if __FILE__ == $0
    TestRunner.parseTest(testLinkedStack())
end
