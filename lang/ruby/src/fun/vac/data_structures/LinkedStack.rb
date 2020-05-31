module LinkedStack

    class Node

        attr_accessor :data
        attr_accessor :next

        public
        def initialize(element)
            @data = element
            @next = nil
        end
    end

    class LinkedStack

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

        def peek()
            if @length.zero?
                raise "[PANIC - NoSuchElement]"
            end

            @tail.data
        end

        def push(element)
            node = Node.new(element)

            if @length.zero?
                @head = node
            else
                @tail.next = node
            end

            @tail = node

            @length += 1
        end

        def pop()
            if @length.zero?
                raise "[PANIC - NoSuchElement]"
            end

            target = @tail

            if @length == 1
                @head = nil
                @tail = nil
            else
                cursor = @head
                (@length - 2).times do
                    cursor = cursor.next
                end
                cursor.next = nil
                @tail = cursor
            end

            target.data = nil

            @length -= 1
        end
    end
end
