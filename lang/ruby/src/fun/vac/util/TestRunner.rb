module TestRunner

    $TestRunnerItemIndex = 0

    def self.parseTest(ok)
        if ok
            puts("✓ Item [%s] PASSED" % $TestRunnerItemIndex)
        else
            STDERR.puts("✗ Item [%s] FAILED" % $TestRunnerItemIndex)
        end

        $TestRunnerItemIndex += 1
    end
end
