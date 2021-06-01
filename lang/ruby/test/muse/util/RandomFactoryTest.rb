require "muse/util/RandomFactory"
require "muse/util/TestRunner"

def testIntegerN()
    RandomFactory.seed()

    value = 0
    8192.times do
        if RandomFactory.integerN(0, 0) != 0
            return false
        end

        if RandomFactory.integerN(1, 1) != 1
            return false
        end

        value = RandomFactory.integerN(0, 1)
        if value.negative? || value > 1
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

    true
end

def testGenerateEven()
    RandomFactory.seed()

    8192.times do
        if (RandomFactory.generateEven() & 1) != 0
            return false
        end
    end

    true
end

def testGenerateOdd()
    RandomFactory.seed()

    8192.times do
        if (RandomFactory.generateOdd() & 1).zero?
            return false
        end
    end

    true
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testIntegerN())

    TestRunner.parseTest(testGenerateEven())

    TestRunner.parseTest(testGenerateOdd())
end
