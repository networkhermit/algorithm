def bruteForce(n)
    arr = Array.new(n, false)

    1.upto(n) do |i|
        (i - 1).step(n - 1, i) do |j|
            arr[j] = !arr[j]
        end
    end

    arr
end

def process(n)
    arr = Array.new(n)

    temp = 0
    1.upto(n) do |i|
        temp = Math.sqrt(i.to_f()).to_i()
        arr[i - 1] = temp * temp == i
    end

    arr
end

if __FILE__ == $PROGRAM_NAME
    N = 100

    arr = process(N)

    arr.each_index do |i|
        if arr[i]
            puts(format("%d\t  OPEN", i + 1))
        end
    end
end
