module SinglyLinkedList
    class Node
        attr_accessor :data
        attr_accessor :next

        public
        def initialize(element)
            @data = element
            @next = nil
        end
    end

    class SinglyLinkedList
        public
        def initialize()
            @head = nil
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

            cursor = @head

            index.times do
                cursor = cursor.next
            end

            cursor.data
        end

        def set(index, element)
            if index.negative? || index >= @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            cursor = @head

            index.times do
                cursor = cursor.next
            end

            cursor.data = element
        end

        def insert(index, element)
            if index.negative? || index > @length
                raise "[PANIC - IndexOutOfBounds]"
            end

            node = Node.new(element)

            if index.zero?
                if @length != 0
                    node.next = @head
                end
                @head = node
            else
                cursor = @head
                (index - 1).times do
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
                else
                    @head = target.next
                end
            else
                cursor = @head
                (index - 1).times do
                    cursor = cursor.next
                end
                target = cursor.next
                cursor.next = target.next
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
