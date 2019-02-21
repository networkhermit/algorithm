require "fun/vac/util/RandomFactory"
require "fun/vac/util/TestRunner"

def testGenerateEven()
    RandomFactory.launch()

    value = 0
    (0 ... 8192).each do |i|
        value = RandomFactory.generateEven()
        if (value & 1) != 0
            return false
        end
    end

    return true
end

def testGenerateOdd()
    RandomFactory.launch()

    value = 0
    (0 ... 8192).each do |i|
        value = RandomFactory.generateOdd()
        if (value & 1) == 0
            return false
        end
    end

    return true
end

if __FILE__ == $0

    TestRunner.parseTest(testGenerateEven())

    TestRunner.parseTest(testGenerateOdd())
end
