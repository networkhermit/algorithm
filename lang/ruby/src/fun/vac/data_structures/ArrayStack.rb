module ArrayStack

    class ArrayStack

        @@DEFAULT_CAPACITY = 64

        public
        def initialize(physicalSize = 0)
            @data = nil
            @logicalSize = 0
            @physicalSize = @@DEFAULT_CAPACITY

            if physicalSize > 1
                @physicalSize = physicalSize
            end
            @data = Array.new(@physicalSize)
        end

        public
        def size()
            @logicalSize
        end

        def isEmpty()
            @logicalSize.zero?
        end

        def peek()
            if @logicalSize.zero?
                raise "[PANIC - NoSuchElement]"
            end

            @data[@logicalSize - 1]
        end

        def push(element)
            if @logicalSize == @physicalSize
                newCapacity = @@DEFAULT_CAPACITY

                if @physicalSize >= @@DEFAULT_CAPACITY
                    newCapacity = @physicalSize + (@physicalSize >> 1)
                end

                temp = Array.new(newCapacity)

                0.upto(@logicalSize - 1) do |i|
                    temp[i] = @data[i]
                end

                @data = temp
                @physicalSize = newCapacity
            end

            @data[@logicalSize] = element

            @logicalSize += 1
        end

        def pop()
            if @logicalSize.zero?
                raise "[PANIC - NoSuchElement]"
            end

            @logicalSize -= 1

            @data[@logicalSize] = nil
        end

        def capacity()
            @physicalSize
        end

        def shrink()
            temp = Array.new(@logicalSize)

            0.upto(@logicalSize - 1) do |i|
                temp[i] = @data[i]
            end

            @data = temp
            @physicalSize = @logicalSize
        end
    end
end
