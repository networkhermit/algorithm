module ArrayList

    class ArrayList

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

        def get(index)
            if index < 0 || index >= @logicalSize
                raise "[PANIC - IndexOutOfBounds]"
            end

            return @data[index]
        end

        def set(index, element)
            if index < 0 || index >= @logicalSize
                raise "[PANIC - IndexOutOfBounds]"
            end

            @data[index] = element
        end

        def insert(index, element)
            if index < 0 || index > @logicalSize
                raise "[PANIC - IndexOutOfBounds]"
            end

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

            i = @logicalSize
            while i > index
                @data[i] = @data[i - 1]
                i -= 1
            end

            @data[index] = element

            @logicalSize += 1
        end

        def remove(index)
            if index < 0 || index >= @logicalSize
                raise "[PANIC - IndexOutOfBounds]"
            end

            for i in index + 1 ... @logicalSize
                @data[i - 1] = @data[i]
            end

            @logicalSize -= 1

            @data[@logicalSize] = nil
        end

        def front()
            return get(0)
        end

        def back()
            return get(@logicalSize - 1)
        end

        def prepend(element)
            insert(0, element)
        end

        def append(element)
            insert(@logicalSize, element)
        end

        def poll()
            remove(0)
        end

        def eject()
            remove(@logicalSize - 1)
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
