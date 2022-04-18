require "muse/data_structures/ArrayStack"
require "muse/util/TestRunner"

def testArrayStack
  size = 8192

  stack = ArrayStack::ArrayStack.new(0)

  1.upto(size) do |i|
    stack.push(i)
  end

  stack.shrink

  return false if stack.size != size

  return false if stack.capacity != size

  size.downto(1) do |i|
    return false if stack.peek != i

    stack.pop
  end

  stack.shrink

  return false unless stack.isEmpty

  stack.capacity.zero?
end

TestRunner.pick(testArrayStack) if __FILE__ == $PROGRAM_NAME
