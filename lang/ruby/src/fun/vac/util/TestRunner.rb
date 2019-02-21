module TestRunner

    $TestRunner_TestIndex = 0

    def self.parseTest(ok)
        if ok
            puts("< Test [%s] Passed >" % $TestRunner_TestIndex)
        else
            STDERR.puts("X Test [%s] Failed X" % $TestRunner_TestIndex)
        end

        $TestRunner_TestIndex += 1
    end
end
