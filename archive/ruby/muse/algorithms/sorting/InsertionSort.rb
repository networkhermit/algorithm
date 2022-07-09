module InsertionSort
  def self.sort(arr)
    target = nil

    cursor = 0

    1.upto(arr.length - 1) do |i|
      target = arr[i]
      cursor = i
      while cursor.positive?
        break if arr[cursor - 1] <= target

        arr[cursor] = arr[cursor - 1]
        cursor -= 1
      end
      arr[cursor] = target
    end
  end
end
