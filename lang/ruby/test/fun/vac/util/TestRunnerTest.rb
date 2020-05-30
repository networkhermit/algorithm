require "fun/vac/util/TestRunner"

def testParseTest()
    10.times do |i|
        TestRunner.parseTest((i & 1) != 0)
    end
end

if __FILE__ == $PROGRAM_NAME
    testParseTest()
end
