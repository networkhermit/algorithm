module PrimeSieves

    def self.sieveOfEratosthenes(n)
        if n < 2
            return Array.new(0)
        end

        size = (n + 1) >> 1

        arr = Array.new(size, false)

        numPrimes = size
        i, bound = 3, Math.sqrt(n).to_i() + 1
        while i < bound
            if arr[i >> 1]
                i += 2
                next
            end
            j = i * i
            while j <= n
                unless arr[j >> 1]
                    arr[j >> 1] = true
                    numPrimes -= 1
                end
                j += i << 1
            end
            i += 2
        end

        primes = Array.new(numPrimes)

        primes[0] = 2

        curr = 1
        i, bound = 3, n + 1
        while i < bound
            unless arr[i >> 1]
                primes[curr] = i
                curr += 1
            end
            i += 2
        end

        primes
    end
end
