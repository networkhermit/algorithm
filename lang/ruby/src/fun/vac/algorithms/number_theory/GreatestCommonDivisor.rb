module GreatestCommonDivisor

    def self.iterativeBinaryGCD(m, n)
        if m.negative?
            m = -m
        end
        if n.negative?
            n = -n
        end

        shift = 0

        loop do
            if m == n
                return m << shift
            end
            if m.zero?
                return n << shift
            end
            if n.zero?
                return m << shift
            end

            if (m & 1).zero?
                m >>= 1
                if (n & 1).zero?
                    n >>= 1
                    shift += 1
                end
            else
                if (n & 1).zero?
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
        if m.negative?
            m = -m
        end
        if n.negative?
            n = -n
        end

        if m == n
            return m
        end
        if m.zero?
            return n
        end
        if n.zero?
            return m
        end

        if (m & 1).zero?
            if (n & 1).zero?
                return recursiveBinaryGCD(m >> 1, n >> 1) << 1
            end
            return recursiveBinaryGCD(m >> 1, n)
        end
        if (n & 1).zero?
            return recursiveBinaryGCD(m, n >> 1)
        end
        if m > n
            return recursiveBinaryGCD((m - n) >> 1, n)
        end
        recursiveBinaryGCD(m, (n - m) >> 1)
    end

    def self.iterativeEuclidean(m, n)
        if m.negative?
            m = -m
        end
        if n.negative?
            n = -n
        end

        until n.zero?
            m, n = n, m % n
        end

        m
    end

    def self.recursiveEuclidean(m, n)
        if m.negative?
            m = -m
        end
        if n.negative?
            n = -n
        end

        if n.zero?
            return m
        end
        recursiveEuclidean(n, m % n)
    end
end
