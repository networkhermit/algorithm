require "fun/vac/util/TestRunner"

def testParseTest()
    for i in 0 ... 10
        TestRunner.parseTest((i & 1) != 0)
    end
end

if __FILE__ == $0
    testParseTest()
end
