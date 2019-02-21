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
            return @logicalSize
        end

        def isEmpty()
            return @logicalSize == 0
        end

        def peek()
            if @logicalSize == 0
                raise "[PANIC - NoSuchElement]"
            end

            return @data[@front]
        end

        def enqueue(element)
            if @logicalSize == @physicalSize
                newCapacity = @@DEFAULT_CAPACITY

                if @physicalSize >= @@DEFAULT_CAPACITY
                    newCapacity = @physicalSize + (@physicalSize >> 1)
                end

                temp = Array.new(newCapacity)

                cursor = @front

                for i in 0 ... @logicalSize
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
            if @logicalSize == 0
                raise "[PANIC - NoSuchElement]"
            end

            @data[@front] = nil

            @front = (@front + 1) % @physicalSize

            @logicalSize -= 1
        end

        def capacity()
            return @physicalSize
        end

        def shrink()
            temp = Array.new(@logicalSize)

            cursor = @front

            for i in 0 ... @logicalSize
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
