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
            @length
        end

        def isEmpty()
            @length.zero?
        end

        def peek()
            if @length.zero?
                raise "[PANIC - NoSuchElement]"
            end

            @head.data
        end

        def enqueue(element)
            node = Node.new(element)

            if @length.zero?
                @head = node
            else
                @tail.next = node
            end

            @tail = node

            @length += 1
        end

        def dequeue()
            if @length.zero?
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
