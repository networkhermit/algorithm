module GreatestCommonDivisor

    def self.iterativeBinaryGCD(m, n)
        if m < 0
            m = -m
        end
        if n < 0
            n = -n
        end

        shift = 0

        while true
            if m == n
                return m << shift
            end
            if m == 0
                return n << shift
            end
            if n == 0
                return m << shift
            end

            if (m & 1) == 0
                m >>= 1
                if (n & 1) == 0
                    n >>= 1
                    shift += 1
                end
            else
                if (n & 1) == 0
                    n >>= 1
                else
                    if m > n
                        m = (m - n) >> 1
                    else
                        n = (n - m) >> 1
                    end
                end
            end
        end
    end

    def self.recursiveBinaryGCD(m, n)
        if m < 0
            m = -m
        end
        if n < 0
            n = -n
        end

        if m == n
            return m
        end
        if m == 0
            return n
        end
        if n == 0
            return m
        end

        if (m & 1) == 0
            if (n & 1) == 0
                return recursiveBinaryGCD(m >> 1, n >> 1) << 1
            end
            return recursiveBinaryGCD(m >> 1, n)
        end
        if (n & 1) == 0
            return recursiveBinaryGCD(m, n >> 1)
        end
        if m > n
            return recursiveBinaryGCD((m - n) >> 1, n)
        end
        return recursiveBinaryGCD(m, (n - m) >> 1)
    end

    def self.iterativeEuclidean(m, n)
        if m < 0
            m = -m
        end
        if n < 0
            n = -n
        end

        until n == 0
            m, n = n, m % n
        end

        return m
    end

    def self.recursiveEuclidean(m, n)
        if m < 0
            m = -m
        end
        if n < 0
            n = -n
        end

        if n == 0
            return m
        end
        return recursiveEuclidean(n, m % n)
    end
end
