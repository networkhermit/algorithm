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
            @logicalSize
        end

        def isEmpty()
            @logicalSize.zero?
        end

        def get(index)
            if index.negative? || index >= @logicalSize
                raise "[PANIC - IndexOutOfBounds]"
            end

            @data[index]
        end

        def set(index, element)
            if index.negative? || index >= @logicalSize
                raise "[PANIC - IndexOutOfBounds]"
            end

            @data[index] = element
        end

        def insert(index, element)
            if index.negative? || index > @logicalSize
                raise "[PANIC - IndexOutOfBounds]"
            end

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

            @logicalSize.downto(index + 1) do |i|
                @data[i] = @data[i - 1]
            end

            @data[index] = element

            @logicalSize += 1
        end

        def remove(index)
            if index.negative? || index >= @logicalSize
                raise "[PANIC - IndexOutOfBounds]"
            end

            (index + 1).upto(@logicalSize - 1) do |i|
                @data[i - 1] = @data[i]
            end

            @logicalSize -= 1

            @data[@logicalSize] = nil
        end

        def front()
            get(0)
        end

        def back()
            get(@logicalSize - 1)
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
