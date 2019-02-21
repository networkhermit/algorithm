module CircularlyLinkedList

    class Node

        attr_accessor :data
        attr_accessor :next

        public
        def initialize(element)
            @data = element
            @next = nil
        end
    end

    class CircularlyLinkedList

        public
        def initialize()
            @tail = nil
            @length = 0
        end

        public
        def size()
            return @length
        end

        def isEmpty()
            return @length == 0
        end

        def get(index)
            if index < 0 || index >= @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            cursor = @tail

            if index != @length - 1
                (0 .. index).each do |i|
                    cursor = cursor.next
                end
            end

            return cursor.data
        end

        def set(index, element)
            if index < 0 || index >= @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            cursor = @tail

            if index != @length - 1
                (0 .. index).each do |i|
                    cursor = cursor.next
                end
            end

            cursor.data = element
        end

        def insert(index, element)
            if index < 0 || index > @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            node = Node.new(element)

            if index == 0
                if @length == 0
                    node.next = node
                    @tail = node
                else
                    node.next = @tail.next
                    @tail.next = node
                end
            elsif index == @length
                node.next = @tail.next
                @tail.next = node
                @tail = node
            else
                cursor = @tail
                (0 .. index - 1).each do |i|
                    cursor = cursor.next
                end
                node.next = cursor.next
                cursor.next = node
            end

            @length += 1
        end

        def remove(index)
            if index < 0 || index >= @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            target = nil

            if index == 0
                target = @tail.next
                if @length == 1
                    @tail = nil
                else
                    @tail.next = target.next
                end
            else
                cursor = @tail
                (0 .. index - 1).each do |i|
                    cursor = cursor.next
                end
                target = cursor.next
                cursor.next = target.next
                if index == @length - 1
                    @tail = cursor
                end
            end

            target.data = nil

            @length -= 1
        end

        def front()
            return get(0)
        end

        def back()
            return get(@length - 1)
        end

        def prepend(element)
            insert(0, element)
        end

        def append(element)
            insert(@length, element)
        end

        def poll()
            remove(0)
        end

        def eject()
            remove(@length - 1)
        end
    end
end
