module TestRunner

    $TestRunnerItemIndex = 0

    def self.parseTest(ok)
        if ok
            puts(format("✓ Item [%s] PASSED", $TestRunnerItemIndex))
        else
            STDERR.puts(format("✗ Item [%s] FAILED", $TestRunnerItemIndex))
        end

        $TestRunnerItemIndex += 1
    end
end
