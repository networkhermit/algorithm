"use strict"

let Node = class {

    constructor(element) {
        this.data = element
        this.next = null
        this.prev = null
    }
}

exports.DoublyLinkedList = class {

    constructor(physicalSize = 0) {
        this.head = null
        this.tail = null
        this.length = 0
    }

    size() {
        return this.length
    }

    isEmpty() {
        return this.length === 0
    }

    get(index) {
        if (index < 0 || index >= this.length) {
            throw new Error("[PANIC - IndexOutOfBounds]")
        }

        let cursor = null

        if (index < this.length >>> 1) {
            cursor = this.head
            for (let i = 0; i < index; i++) {
                cursor = cursor.next
            }
        } else {
            cursor = this.tail
            for (let i = this.length - 1; i > index; i--) {
                cursor = cursor.prev
            }
        }

        return cursor.data
    }

    set(index, element) {
        if (index < 0 || index >= this.length) {
            throw new Error("[PANIC - IndexOutOfBounds]")
        }

        let cursor = null

        if (index < this.length >>> 1) {
            cursor = this.head
            for (let i = 0; i < index; i++) {
                cursor = cursor.next
            }
        } else {
            cursor = this.tail
            for (let i = this.length - 1; i > index; i--) {
                cursor = cursor.prev
            }
        }

        cursor.data = element
    }

    insert(index, element) {
        if (index < 0 || index > this.length) {
            throw new Error("[PANIC - IndexOutOfBounds]")
        }

        let node = new Node(element)

        if (index === 0) {
            if (this.length !== 0) {
                node.next = this.head
                this.head.prev = node
            } else {
                this.tail = node
            }
            this.head = node
        } else if (index === this.length) {
            node.prev = this.tail
            this.tail.next = node
            this.tail = node
        } else {
            let cursor = null
            if (index < this.length >>> 1) {
                cursor = this.head
                for (let i = 0; i < index; i++) {
                    cursor = cursor.next
                }
            } else {
                cursor = this.tail
                for (let i = this.length - 1; i > index; i--) {
                    cursor = cursor.prev
                }
            }
            node.next = cursor
            node.prev = cursor.prev
            cursor.prev.next = node
            cursor.prev = node
        }

        this.length += 1
    }

    remove(index) {
        if (index < 0 || index >= this.length) {
            throw new Error("[PANIC - IndexOutOfBounds]")
        }

        let target = null

        if (index === 0) {
            target = this.head
            if (this.length === 1) {
                this.head = null
                this.tail = null
            } else {
                target.next.prev = null
                this.head = target.next
            }
        } else if (index === this.length - 1) {
            target = this.tail
            target.prev.next = null
            this.tail = target.prev
        } else {
            if (index < this.length >>> 1) {
                target = this.head
                for (let i = 0; i < index; i++) {
                    target = target.next
                }
            } else {
                target = this.tail
                for (let i = this.length - 1; i > index; i--) {
                    target = target.prev
                }
            }
            target.prev.next = target.next
            target.next.prev = target.prev
        }

        target.data = null

        this.length -= 1
    }

    front() {
        return this.get(0)
    }

    back() {
        return this.get(this.length - 1)
    }

    prepend(element) {
        this.insert(0, element)
    }

    append(element) {
        this.insert(this.length, element)
    }

    poll() {
        this.remove(0)
    }

    eject() {
        this.remove(this.length - 1)
    }
}
