module LinkedQueue

    class Node

        attr_accessor :data
        attr_accessor :next

        public
        def initialize(element)
            @data = element
            @next = nil
        end
    end

    class LinkedQueue

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

        def peek()
            if @length == 0
                raise "[PANIC - NoSuchElement]"
            end

            return @head.data
        end

        def enqueue(element)
            node = Node.new(element)

            if @length == 0
                @head = node
            else
                @tail.next = node
            end

            @tail = node

            @length += 1
        end

        def dequeue()
            if @length == 0
                raise "[PANIC - NoSuchElement]"
            end

            target = @head

            if @length == 1
                @head = nil
                @tail = nil
            else
                @head = target.next
            end

            target.data = nil

            @length -= 1
        end
    end
end
