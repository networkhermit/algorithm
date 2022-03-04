def fizzBuzz(n)
    word = ""

    if (n % 3).zero?
        word += "Fizz"
    end
    if (n % 5).zero?
        word += "Buzz"
    end

    if word.empty?
        word = n.to_s()
    end

    word
end

if __FILE__ == $PROGRAM_NAME
    1.upto(100) do |i|
        puts(fizzBuzz(i))
    end
end
