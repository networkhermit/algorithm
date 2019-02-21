module InsertionSort

    def self.sort(arr)
        target = nil

        cursor = 0

        for i in 1 ... arr.length
            target = arr[i]
            cursor = i
            while cursor > 0
                if arr[cursor - 1] > target
                    arr[cursor] = arr[cursor - 1]
                else
                    break
                end
                cursor -= 1
            end
            arr[cursor] = target
        end
    end
end
