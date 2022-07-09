module SinglyLinkedList
  class Node
    attr_accessor :data, :next

    def initialize(element)
      @data = element
      @next = nil
    end
  end

  class SinglyLinkedList
    def initialize
      @head = nil
      @length = 0
    end

    def size
      @length
    end

    def isEmpty
      @length.zero?
    end

    def get(index)
      raise "[PANIC - IndexOutOfBounds]" if index.negative? || index >= @length

      cursor = @head

      index.times do
        cursor = cursor.next
      end

      cursor.data
    end

    def set(index, element)
      raise "[PANIC - IndexOutOfBounds]" if index.negative? || index >= @length

      cursor = @head

      index.times do
        cursor = cursor.next
      end

      cursor.data = element
    end

    def insert(index, element)
      raise "[PANIC - IndexOutOfBounds]" if index.negative? || index > @length

      node = Node.new(element)

      if index.zero?
        node.next = @head if @length != 0
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
      raise "[PANIC - IndexOutOfBounds]" if index.negative? || index >= @length

      target = nil

      if index.zero?
        target = @head
        @head = if @length == 1
          nil
        else
          target.next
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

    def front
      get(0)
    end

    def back
      get(@length - 1)
    end

    def prepend(element)
      insert(0, element)
    end

    def append(element)
      insert(@length, element)
    end

    def poll
      remove(0)
    end

    def eject
      remove(@length - 1)
    end
  end
end
