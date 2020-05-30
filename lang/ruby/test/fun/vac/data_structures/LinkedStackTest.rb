require "fun/vac/data_structures/LinkedStack"
require "fun/vac/util/TestRunner"

def testLinkedStack()
    size = 8192

    stack = LinkedStack::LinkedStack.new()

    (1 .. size).each do |i|
        stack.push(i)
    end

    if stack.size() != size
        return false
    end

    i = size
    while i.positive?
        if stack.peek() != i
            return false
        end
        stack.pop()
        i -= 1
    end

    stack.isEmpty()
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testLinkedStack())
end
