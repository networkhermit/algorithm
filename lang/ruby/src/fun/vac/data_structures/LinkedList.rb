module LinkedList

    class Node

        attr_accessor :data
        attr_accessor :next

        public
        def initialize(element)
            @data = element
            @next = nil
        end
    end

    class LinkedList

        public
        def initialize()
            @head = nil
            @tail = nil
            @length = 0
        end

        public
        def size()
            @length
        end

        def isEmpty()
            @length.zero?
        end

        def get(index)
            if index.negative? || index >= @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            cursor = nil

            if index == @length - 1
                cursor = @tail
            else
                cursor = @head
                (0 ... index).each do
                    cursor = cursor.next
                end
            end

            cursor.data
        end

        def set(index, element)
            if index.negative? || index >= @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            cursor = nil

            if index == @length - 1
                cursor = @tail
            else
                cursor = @head
                (0 ... index).each do
                    cursor = cursor.next
                end
            end

            cursor.data = element
        end

        def insert(index, element)
            if index.negative? || index > @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            node = Node.new(element)

            case index
            when 0
                if @length != 0
                    node.next = @head
                else
                    @tail = node
                end
                @head = node
            when @length
                @tail.next = node
                @tail = node
            else
                cursor = @head
                (0 ... index - 1).each do
                    cursor = cursor.next
                end
                node.next = cursor.next
                cursor.next = node
            end

            @length += 1
        end

        def remove(index)
            if index.negative? || index >= @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            target = nil

            if index.zero?
                target = @head
                if @length == 1
                    @head = nil
                    @tail = nil
                else
                    @head = target.next
                end
            else
                cursor = @head
                (0 ... index - 1).each do
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
            get(0)
        end

        def back()
            get(@length - 1)
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
