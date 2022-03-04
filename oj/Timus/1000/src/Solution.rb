if __FILE__ == $PROGRAM_NAME
    while (line = gets())
        pair = line.split()
        a = pair[0].to_i()
        b = pair[1].to_i()
        puts(a + b)
    end
end
