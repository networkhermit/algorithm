require "fun/vac/util/TestRunner"

def testParseTest()
    for i in 0 ... 10
        if (i & 1) == 0
            TestRunner.parseTest(false)
        else
            TestRunner.parseTest(true)
        end
    end
end

if __FILE__ == $0
    testParseTest()
end
