require "fun/vac/util/RandomFactory"
require "fun/vac/util/TestRunner"

def testIntegerN()
    RandomFactory.seed()

    value = 0
    (0 ... 8192).each do |i|
        if RandomFactory.integerN(0, 0) != 0
            return false
        end

        if RandomFactory.integerN(1, 1) != 1
            return false
        end

        value = RandomFactory.integerN(0, 1)
        if value < 0 || value > 1
            return false
        end

        value = RandomFactory.integerN(100, 10000)
        if RandomFactory.integerN(value, value) != value
            return false
        end
        if value < 100 || value > 10000
            return false
        end
    end

    return true
end

def testGenerateEven()
    RandomFactory.seed()

    (0 ... 8192).each do |i|
        if (RandomFactory.generateEven() & 1) != 0
            return false
        end
    end

    return true
end

def testGenerateOdd()
    RandomFactory.seed()

    (0 ... 8192).each do |i|
        if (RandomFactory.generateOdd() & 1) == 0
            return false
        end
    end

    return true
end

if __FILE__ == $0

    TestRunner.parseTest(testIntegerN())

    TestRunner.parseTest(testGenerateEven())

    TestRunner.parseTest(testGenerateOdd())
end
