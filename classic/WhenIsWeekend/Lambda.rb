if __FILE__ == $PROGRAM_NAME
    puts("When’s Weekend?")

    today = Time.now().wday

    case today
    when 0
        puts("Today.")
    when 6
        puts("Today.")
    when 5
        puts("Tomorrow.")
    when 4
        puts("In two days.")
    else
        puts("Too far away.")
    end
end
