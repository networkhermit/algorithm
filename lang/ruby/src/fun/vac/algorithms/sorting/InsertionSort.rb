module InsertionSort

    def self.sort(arr)
        target = nil

        cursor = 0

        (1 ... arr.length).each do |i|
            target = arr[i]
            cursor = i
            while cursor.positive?
                if arr[cursor - 1] <= target
                    break
                end
                arr[cursor] = arr[cursor - 1]
                cursor -= 1
            end
            arr[cursor] = target
        end
    end
end
