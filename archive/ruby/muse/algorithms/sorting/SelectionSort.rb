module SelectionSort
  def self.sort(arr)
    return if arr.empty?

    iMin = 0

    bound = arr.length - 1
    0.upto(bound - 1) do |i|
      iMin = i
      (i + 1).upto(bound) do |j|
        iMin = j if arr[j] < arr[iMin]
      end
      arr[i], arr[iMin] = arr[iMin], arr[i] if iMin != i
    end
  end
end
