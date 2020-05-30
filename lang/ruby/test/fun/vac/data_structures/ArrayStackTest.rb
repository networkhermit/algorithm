require "fun/vac/data_structures/ArrayStack"
require "fun/vac/util/TestRunner"

def testArrayStack()
    size = 8192

    stack = ArrayStack::ArrayStack.new(0)

    (1 .. size).each do |i|
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
    while i.positive?
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

    stack.capacity().zero?
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testArrayStack())
end
