module SelectionSort

    def self.sort(arr)
        if arr.empty?
            return
        end

        iMin = 0

        bound = arr.length - 1
        0.upto(bound - 1) do |i|
            iMin = i
            (i + 1).upto(bound) do |j|
                if arr[j] < arr[iMin]
                    iMin = j
                end
            end
            if iMin != i
                arr[i], arr[iMin] = arr[iMin], arr[i]
            end
        end
    end
end
