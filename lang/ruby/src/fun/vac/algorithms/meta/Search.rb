require "fun/vac/algorithms/search/BinarySearch"
require "fun/vac/algorithms/search/LinearSearch"

module Search

    def self.binarySearch(arr, key)
        return BinarySearch.find(arr, key)
    end

    def self.linearSearch(arr, key)
        return LinearSearch.find(arr, key)
    end
end
