require "fun/vac/algorithms/sorting/BubbleSort"
require "fun/vac/algorithms/sorting/InsertionSort"
require "fun/vac/algorithms/sorting/MergeSort"
require "fun/vac/algorithms/sorting/QuickSort"
require "fun/vac/algorithms/sorting/SelectionSort"

module Sorting

    def self.bubbleSort(arr)
        BubbleSort.sort(arr)
    end

    def self.insertionSort(arr)
        InsertionSort.sort(arr)
    end

    def self.mergeSort(arr)
        MergeSort.sort(arr)
    end

    def self.quickSort(arr)
        QuickSort.sort(arr)
    end

    def self.selectionSort(arr)
        SelectionSort.sort(arr)
    end
end
