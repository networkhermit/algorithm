require "vac.fun/algorithms/search/BinarySearch"
require "vac.fun/algorithms/search/LinearSearch"

module Search

    def self.binarySearch(arr, key)
        BinarySearch.find(arr, key)
    end

    def self.linearSearch(arr, key)
        LinearSearch.find(arr, key)
    end
end
