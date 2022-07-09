module LinkedStack
  class Node
    attr_accessor :data, :next

    def initialize(element)
      @data = element
      @next = nil
    end
  end

  class LinkedStack
    def initialize
      @head = nil
      @tail = nil
      @length = 0
    end

    def size
      @length
    end

    def isEmpty
      @length.zero?
    end

    def peek
      raise "[PANIC - NoSuchElement]" if @length.zero?

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

    def pop
      raise "[PANIC - NoSuchElement]" if @length.zero?

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
