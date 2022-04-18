require "muse/util/RandomFactory"
require "muse/util/TestRunner"

def testGenIntN
  value = 0
  8192.times do
    return false if RandomFactory.genIntN(0, 0) != 0

    return false if RandomFactory.genIntN(1, 1) != 1

    value = RandomFactory.genIntN(0, 1)
    return false if value.negative? || value > 1

    value = RandomFactory.genIntN(100, 10_000)
    return false if RandomFactory.genIntN(value, value) != value
    return false if value < 100 || value > 10_000
  end

  true
end

def testGenEven
  8192.times do
    return false if (RandomFactory.genEven & 1) != 0
  end

  true
end

def testGenOdd
  8192.times do
    return false if (RandomFactory.genOdd & 1).zero?
  end

  true
end

if __FILE__ == $PROGRAM_NAME
  TestRunner.pick(testGenIntN)

  TestRunner.pick(testGenEven)

  TestRunner.pick(testGenOdd)
end
