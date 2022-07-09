require "securerandom"

module RandomFactory
  def self.genIntN(min, max)
    min + SecureRandom.random_number(max - min + 1)
  end

  def self.genInt
    genIntN(1, 2_147_483_647)
  end

  def self.genEven
    genInt >> 1 << 1
  end

  def self.genOdd
    genInt | 1
  end
end
