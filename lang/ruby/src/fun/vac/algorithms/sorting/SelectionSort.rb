module SelectionSort

    def self.sort(arr)
        if arr.length == 0
            return
        end

        iMin = 0

        bound = arr.length - 1
        for i in 0 ... bound
            iMin = i
            for j in i + 1 .. bound
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
