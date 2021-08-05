require "muse/util/TestRunner"

def testPick()
    0.upto(9) do |i|
        TestRunner.pick((i & 1) != 0)
    end
end

if __FILE__ == $PROGRAM_NAME
    testPick()
end
