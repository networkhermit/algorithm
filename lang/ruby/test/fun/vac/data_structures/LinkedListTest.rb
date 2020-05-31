require "fun/vac/data_structures/LinkedList"
require "fun/vac/util/TestRunner"

def testLinkedList()
    size = 8192

    list = LinkedList::LinkedList.new()

    1.upto(size) do |i|
        list.append(i)
    end

    if list.size() != size
        return false
    end

    0.upto(size - 1) do |i|
        if list.get(i) != i + 1
            return false
        end
    end

    0.upto(size - 1) do |i|
        list.set(i, size - i)
    end

    0.upto(size - 1) do |i|
        if list.get(i) != size - i
            return false
        end
    end

    x = 0
    y = 0

    i, j = 0, size - 1
    while i < j
        x = list.get(i)
        y = list.get(j)

        list.remove(i)
        list.insert(i, y)
        list.remove(j)
        list.insert(j, x)
        i, j = i + 1, j - 1
    end

    0.upto(size - 1) do |i|
        if list.get(i) != i + 1
            return false
        end
    end

    size.downto(1) do |i|
        if list.back() != i
            return false
        end
        list.eject()
    end

    list.isEmpty()
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testLinkedList())
end
