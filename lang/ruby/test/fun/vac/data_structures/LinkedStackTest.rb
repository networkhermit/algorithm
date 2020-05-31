require "fun/vac/data_structures/LinkedStack"
require "fun/vac/util/TestRunner"

def testLinkedStack()
    size = 8192

    stack = LinkedStack::LinkedStack.new()

    1.upto(size) do |i|
        stack.push(i)
    end

    if stack.size() != size
        return false
    end

    size.downto(1) do |i|
        if stack.peek() != i
            return false
        end
        stack.pop()
    end

    stack.isEmpty()
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testLinkedStack())
end
