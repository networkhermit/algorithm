module PrimeSieves

    def self.sieveOfEratosthenes(n)
        if n < 2
            return Array.new(0)
        end

        size = (n + 1) >> 1

        arr = Array.new(size, false)

        numPrimes = size
        3.step(Math.sqrt(n).to_i(), 2) do |i|
            if arr[i >> 1]
                next
            end
            (i * i).step(n, i << 1) do |j|
                unless arr[j >> 1]
                    arr[j >> 1] = true
                    numPrimes -= 1
                end
            end
        end

        primes = Array.new(numPrimes)

        primes[0] = 2

        curr = 1
        3.step(n, 2) do |i|
            unless arr[i >> 1]
                primes[curr] = i
                curr += 1
            end
        end

        primes
    end
end
