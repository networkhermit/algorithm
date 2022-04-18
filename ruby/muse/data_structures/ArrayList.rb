module ArrayList
  class ArrayList
    @@DEFAULT_CAPACITY = 64

    def initialize(physicalSize = 0)
      @data = nil
      @logicalSize = 0
      @physicalSize = @@DEFAULT_CAPACITY

      @physicalSize = physicalSize if physicalSize > 1
      @data = Array.new(@physicalSize)
    end

    def size
      @logicalSize
    end

    def isEmpty
      @logicalSize.zero?
    end

    def get(index)
      raise "[PANIC - IndexOutOfBounds]" if index.negative? || index >= @logicalSize

      @data[index]
    end

    def set(index, element)
      raise "[PANIC - IndexOutOfBounds]" if index.negative? || index >= @logicalSize

      @data[index] = element
    end

    def insert(index, element)
      raise "[PANIC - IndexOutOfBounds]" if index.negative? || index > @logicalSize

      if @logicalSize == @physicalSize
        newCapacity = @@DEFAULT_CAPACITY

        newCapacity = @physicalSize + (@physicalSize >> 1) if @physicalSize >= @@DEFAULT_CAPACITY

        temp = Array.new(newCapacity)

        0.upto(@logicalSize - 1) do |i|
          temp[i] = @data[i]
        end

        @data = temp
        @physicalSize = newCapacity
      end

      @logicalSize.downto(index + 1) do |i|
        @data[i] = @data[i - 1]
      end

      @data[index] = element

      @logicalSize += 1
    end

    def remove(index)
      raise "[PANIC - IndexOutOfBounds]" if index.negative? || index >= @logicalSize

      (index + 1).upto(@logicalSize - 1) do |i|
        @data[i - 1] = @data[i]
      end

      @logicalSize -= 1

      @data[@logicalSize] = nil
    end

    def front
      get(0)
    end

    def back
      get(@logicalSize - 1)
    end

    def prepend(element)
      insert(0, element)
    end

    def append(element)
      insert(@logicalSize, element)
    end

    def poll
      remove(0)
    end

    def eject
      remove(@logicalSize - 1)
    end

    def capacity
      @physicalSize
    end

    def shrink
      temp = Array.new(@logicalSize)

      0.upto(@logicalSize - 1) do |i|
        temp[i] = @data[i]
      end

      @data = temp
      @physicalSize = @logicalSize
    end
  end
end
