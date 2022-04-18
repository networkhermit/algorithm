def fizzBuzz(n)
  word = ""

  word += "Fizz" if (n % 3).zero?
  word += "Buzz" if (n % 5).zero?

  word = n.to_s if word.empty?

  word
end

if __FILE__ == $PROGRAM_NAME
  1.upto(100) do |i|
    puts(fizzBuzz(i))
  end
end
