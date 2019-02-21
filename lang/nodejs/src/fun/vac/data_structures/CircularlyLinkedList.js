"use strict"

let Node = class {

    constructor(element) {
        this.data = element
        this.next = null
    }
}

exports.CircularlyLinkedList = class {

    constructor(physicalSize = 0) {
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

        let cursor = this.tail

        if (index !== this.length - 1) {
            for (let i = 0; i <= index; i++) {
                cursor = cursor.next
            }
        }

        return cursor.data
    }

    set(index, element) {
        if (index < 0 || index >= this.length) {
            throw new Error("[PANIC - IndexOutOfBounds]")
        }

        let cursor = this.tail

        if (index !== this.length - 1) {
            for (let i = 0; i <= index; i++) {
                cursor = cursor.next
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
            if (this.length === 0) {
                node.next = node
                this.tail = node
            } else {
                node.next = this.tail.next
                this.tail.next = node
            }
        } else if (index === this.length) {
            node.next = this.tail.next
            this.tail.next = node
            this.tail = node
        } else {
            let cursor = this.tail
            for (let i = 0, bound = index - 1; i <= bound; i++) {
                cursor = cursor.next
            }
            node.next = cursor.next
            cursor.next = node
        }

        this.length += 1
    }

    remove(index) {
        if (index < 0 || index >= this.length) {
            throw new Error("[PANIC - IndexOutOfBounds]")
        }

        let target = null

        if (index === 0) {
            target = this.tail.next
            if (this.length === 1) {
                this.tail = null
            } else {
                this.tail.next = target.next
            }
        } else {
            let cursor = this.tail
            for (let i = 0, bound = index - 1; i <= bound; i++) {
                cursor = cursor.next
            }
            target = cursor.next
            cursor.next = target.next
            if (index === this.length - 1) {
                this.tail = cursor
            }
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
