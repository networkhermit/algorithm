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
            return @logicalSize
        end

        def isEmpty()
            return @logicalSize == 0
        end

        def peek()
            if @logicalSize == 0
                raise "[PANIC - NoSuchElement]"
            end

            return @data[@logicalSize - 1]
        end

        def push(element)
            if @logicalSize == @physicalSize
                newCapacity = @@DEFAULT_CAPACITY

                if @physicalSize >= @@DEFAULT_CAPACITY
                    newCapacity = @physicalSize + (@physicalSize >> 1)
                end

                temp = Array.new(newCapacity)

                for i in 0 ... @logicalSize
                    temp[i] = @data[i]
                end

                @data = temp
                @physicalSize = newCapacity
            end

            @data[@logicalSize] = element

            @logicalSize += 1
        end

        def pop()
            if @logicalSize == 0
                raise "[PANIC - NoSuchElement]"
            end

            @logicalSize -= 1

            @data[@logicalSize] = nil
        end

        def capacity()
            return @physicalSize
        end

        def shrink()
            temp = Array.new(@logicalSize)

            for i in 0 ... @logicalSize
                temp[i] = @data[i]
            end

            @data = temp
            @physicalSize = @logicalSize
        end
    end
end
