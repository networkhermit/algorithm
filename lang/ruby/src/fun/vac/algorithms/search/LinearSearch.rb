module LinearSearch

    def self.find(arr, key)
        arr.each_index do |i|
            if arr[i] == key
                return i
            end
        end

        return arr.length
    end
end
