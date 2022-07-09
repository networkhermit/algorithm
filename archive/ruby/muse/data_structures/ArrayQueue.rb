module ArrayQueue
  class ArrayQueue
    @@DEFAULT_CAPACITY = 64

    def initialize(physicalSize = 0)
      @data = nil
      @front = 0
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

    def peek
      raise "[PANIC - NoSuchElement]" if @logicalSize.zero?

      @data[@front]
    end

    def enqueue(element)
      if @logicalSize == @physicalSize
        newCapacity = @@DEFAULT_CAPACITY

        newCapacity = @physicalSize + (@physicalSize >> 1) if @physicalSize >= @@DEFAULT_CAPACITY

        temp = Array.new(newCapacity)

        cursor = @front

        0.upto(@logicalSize - 1) do |i|
          cursor = 0 if cursor == @physicalSize
          temp[i] = @data[cursor]
          cursor += 1
        end

        @data = temp
        @front = 0
        @physicalSize = newCapacity
      end

      @data[(@front + @logicalSize) % @physicalSize] = element

      @logicalSize += 1
    end

    def dequeue
      raise "[PANIC - NoSuchElement]" if @logicalSize.zero?

      @data[@front] = nil

      @front = (@front + 1) % @physicalSize

      @logicalSize -= 1
    end

    def capacity
      @physicalSize
    end

    def shrink
      temp = Array.new(@logicalSize)

      cursor = @front

      0.upto(@logicalSize - 1) do |i|
        cursor = 0 if cursor == @physicalSize
        temp[i] = @data[cursor]
        cursor += 1
      end

      @data = temp
      @front = 0
      @physicalSize = @logicalSize
    end
  end
end
