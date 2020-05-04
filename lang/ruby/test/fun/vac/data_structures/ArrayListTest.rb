require "fun/vac/data_structures/ArrayList"
require "fun/vac/util/TestRunner"

def testArrayList()
    size = 8192

    list = ArrayList::ArrayList.new(0)

    for i in 1 .. size
        list.append(i)
    end

    list.shrink()

    if list.size() != size
        return false
    end

    if list.capacity() != size
        return false
    end

    for i in 0 ... size
        if list.get(i) != i + 1
            return false
        end
    end

    for i in 0 ... size
        list.set(i, size - i)
    end

    for i in 0 ... size
        if list.get(i) != size - i
            return false
        end
    end

    x, y = 0, 0

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

    for i in 0 ... size
        if list.get(i) != i + 1
            return false
        end
    end

    i = size
    while i >= 1
        if list.back() != i
            return false
        end
        list.eject()
        i -= 1
    end

    list.shrink()

    unless list.isEmpty()
        return false
    end

    if list.capacity() != 0
        return false
    end

    return true
end

if __FILE__ == $0
    TestRunner.parseTest(testArrayList())
end
