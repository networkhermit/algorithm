require "muse/algorithms/sorting/BubbleSort"
require "muse/algorithms/sorting/InsertionSort"
require "muse/algorithms/sorting/MergeSort"
require "muse/algorithms/sorting/QuickSort"
require "muse/algorithms/sorting/SelectionSort"

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
