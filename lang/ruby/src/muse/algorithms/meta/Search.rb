require "muse/algorithms/search/BinarySearch"
require "muse/algorithms/search/LinearSearch"

module Search

    def self.binarySearch(arr, key)
        BinarySearch.find(arr, key)
    end

    def self.linearSearch(arr, key)
        LinearSearch.find(arr, key)
    end
end
