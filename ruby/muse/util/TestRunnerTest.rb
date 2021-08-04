require "muse/util/TestRunner"

def testParseTest()
    0.upto(9) do |i|
        TestRunner.parseTest((i & 1) != 0)
    end
end

if __FILE__ == $PROGRAM_NAME
    testParseTest()
end
