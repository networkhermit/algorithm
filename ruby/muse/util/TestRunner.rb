module TestRunner

    @@itemIndex = 0

    def self.pick(ok)
        if ok
            puts(format("✓ Item [%s] PASSED", @@itemIndex))
        else
            STDERR.puts(format("✗ Item [%s] FAILED", @@itemIndex))
        end

        @@itemIndex += 1
    end
end
