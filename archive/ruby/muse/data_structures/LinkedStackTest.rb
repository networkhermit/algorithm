require "muse/data_structures/LinkedStack"
require "muse/util/TestRunner"

def testLinkedStack
  size = 8192

  stack = LinkedStack::LinkedStack.new

  1.upto(size) do |i|
    stack.push(i)
  end

  return false if stack.size != size

  size.downto(1) do |i|
    return false if stack.peek != i

    stack.pop
  end

  stack.isEmpty
end

TestRunner.pick(testLinkedStack) if __FILE__ == $PROGRAM_NAME
