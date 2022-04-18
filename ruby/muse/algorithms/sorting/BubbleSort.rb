module BubbleSort
  def self.sort(arr)
    margin = 0
    unsorted = arr.length

    while unsorted > 1
      margin = 0
      1.upto(unsorted - 1) do |i|
        if arr[i - 1] > arr[i]
          arr[i - 1], arr[i] = arr[i], arr[i - 1]
          margin = i
        end
      end
      unsorted = margin
    end
  end
end
