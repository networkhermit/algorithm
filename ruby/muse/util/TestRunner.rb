module TestRunner
  @@itemIndex = 0

  def self.pick(ok)
    if ok
      puts(format("✓ Item [%s] PASSED", @@itemIndex))
    else
      warn(format("✗ Item [%s] FAILED", @@itemIndex))
    end

    @@itemIndex += 1
  end
end
