module LinkedQueue
  class Node
    attr_accessor :data, :next

    def initialize(element)
      @data = element
      @next = nil
    end
  end

  class LinkedQueue
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

    def dequeue
      raise "[PANIC - NoSuchElement]" if @length.zero?

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
