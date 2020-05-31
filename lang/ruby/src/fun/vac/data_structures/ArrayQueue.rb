module ArrayQueue

    class ArrayQueue

        @@DEFAULT_CAPACITY = 64

        public
        def initialize(physicalSize = 0)
            @data = nil
            @front = 0
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

            @data[@front]
        end

        def enqueue(element)
            if @logicalSize == @physicalSize
                newCapacity = @@DEFAULT_CAPACITY

                if @physicalSize >= @@DEFAULT_CAPACITY
                    newCapacity = @physicalSize + (@physicalSize >> 1)
                end

                temp = Array.new(newCapacity)

                cursor = @front

                0.upto(@logicalSize - 1) do |i|
                    if cursor == @physicalSize
                        cursor = 0
                    end
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

        def dequeue()
            if @logicalSize.zero?
                raise "[PANIC - NoSuchElement]"
            end

            @data[@front] = nil

            @front = (@front + 1) % @physicalSize

            @logicalSize -= 1
        end

        def capacity()
            @physicalSize
        end

        def shrink()
            temp = Array.new(@logicalSize)

            cursor = @front

            0.upto(@logicalSize - 1) do |i|
                if cursor == @physicalSize
                    cursor = 0
                end
                temp[i] = @data[cursor]
                cursor += 1
            end

            @data = temp
            @front = 0
            @physicalSize = @logicalSize
        end
    end
end
