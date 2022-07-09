require "muse/data_structures/DoublyLinkedList"
require "muse/util/TestRunner"

def testLinkedList
  size = 8192

  list = DoublyLinkedList::DoublyLinkedList.new

  1.upto(size) do |i|
    list.append(i)
  end

  return false if list.size != size

  0.upto(size - 1) do |i|
    return false if list.get(i) != i + 1
  end

  0.upto(size - 1) do |i|
    list.set(i, size - i)
  end

  0.upto(size - 1) do |i|
    return false if list.get(i) != size - i
  end

  i = 0
  j = size - 1
  while i < j
    x = list.get(i)
    y = list.get(j)

    list.remove(i)
    list.insert(i, y)
    list.remove(j)
    list.insert(j, x)
    i += 1
    j -= 1
  end

  0.upto(size - 1) do |i|
    return false if list.get(i) != i + 1
  end

  size.downto(1) do |i|
    return false if list.back != i

    list.eject
  end

  list.isEmpty
end

TestRunner.pick(testLinkedList) if __FILE__ == $PROGRAM_NAME
