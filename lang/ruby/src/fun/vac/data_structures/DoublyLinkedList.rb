module DoublyLinkedList

    class Node

        attr_accessor :data
        attr_accessor :next
        attr_accessor :prev

        public
        def initialize(element)
            @data = element
            @next = nil
            @prev = nil
        end
    end

    class DoublyLinkedList

        public
        def initialize()
            @head = nil
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

            cursor = nil

            if index < @length >> 1
                cursor = @head
                (0 ... index).each do |i|
                    cursor = cursor.next
                end
            else
                cursor = @tail
                i = @length - 1
                while i > index
                    cursor = cursor.prev
                    i -= 1
                end
            end

            return cursor.data
        end

        def set(index, element)
            if index < 0 || index >= @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            cursor = nil

            if index < @length >> 1
                cursor = @head
                (0 ... index).each do |i|
                    cursor = cursor.next
                end
            else
                cursor = @tail
                i = @length - 1
                while i > index
                    cursor = cursor.prev
                    i -= 1
                end
            end

            cursor.data = element
        end

        def insert(index, element)
            if index < 0 || index > @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            node = Node.new(element)

            case index
            when 0
                if @length != 0
                    node.next = @head
                    @head.prev = node
                else
                    @tail = node
                end
                @head = node
            when @length
                node.prev = @tail
                @tail.next = node
                @tail = node
            else
                cursor = nil
                if index < @length >> 1
                    cursor = @head
                    (0 ... index).each do |i|
                        cursor = cursor.next
                    end
                else
                    cursor = @tail
                    i = @length - 1
                    while i > index
                        cursor = cursor.prev
                        i -= 1
                    end
                end
                node.next = cursor
                node.prev = cursor.prev
                cursor.prev.next = node
                cursor.prev = node
            end

            @length += 1
        end

        def remove(index)
            if index < 0 || index >= @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            target = nil

            case index
            when 0
                target = @head
                if @length == 1
                    @head = nil
                    @tail = nil
                else
                    target.next.prev = nil
                    @head = target.next
                end
            when @length - 1
                target = @tail
                target.prev.next = nil
                @tail = target.prev
            else
                if index < @length >> 1
                    target = @head
                    (0 ... index).each do |i|
                        target = target.next
                    end
                else
                    target = @tail
                    i = @length - 1
                    while i > index
                        target = target.prev
                        i -= 1
                    end
                end
                target.prev.next = target.next
                target.next.prev = target.prev
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
