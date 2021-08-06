require "muse/util/RandomFactory"
require "muse/util/TestRunner"

def testGenIntN()
    value = 0
    8192.times do
        if RandomFactory.genIntN(0, 0) != 0
            return false
        end

        if RandomFactory.genIntN(1, 1) != 1
            return false
        end

        value = RandomFactory.genIntN(0, 1)
        if value.negative? || value > 1
            return false
        end

        value = RandomFactory.genIntN(100, 10000)
        if RandomFactory.genIntN(value, value) != value
            return false
        end
        if value < 100 || value > 10000
            return false
        end
    end

    true
end

def testGenEven()
    8192.times do
        if (RandomFactory.genEven() & 1) != 0
            return false
        end
    end

    true
end

def testGenOdd()
    8192.times do
        if (RandomFactory.genOdd() & 1).zero?
            return false
        end
    end

    true
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.pick(testGenIntN())

    TestRunner.pick(testGenEven())

    TestRunner.pick(testGenOdd())
end
