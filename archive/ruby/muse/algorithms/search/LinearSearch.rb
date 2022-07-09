module LinearSearch
  def self.find(arr, key)
    arr.each_index do |i|
      return i if arr[i] == key
    end

    arr.length
  end
end
