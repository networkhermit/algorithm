"use strict"

let Node = class {

    constructor(element) {
        this.data = element
        this.next = null
    }
}

exports.LinkedStack = class {

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

    peek() {
        if (this.length === 0) {
            throw new Error("[PANIC - NoSuchElement]")
        }

        return this.tail.data
    }

    push(element) {
        let node = new Node(element)

        if (this.length === 0) {
            this.head = node
        } else {
            this.tail.next = node
        }

        this.tail = node

        this.length++
    }

    pop() {
        if (this.length === 0) {
            throw new Error("[PANIC - NoSuchElement]")
        }

        let target = this.tail

        if (this.length === 1) {
            this.head = null
            this.tail = null
        } else {
            let cursor = this.head
            for (let i = 0, bound = this.length - 2; i < bound; i++) {
                cursor = cursor.next
            }
            cursor.next = null
            this.tail = cursor
        }

        target.data = null

        this.length--
    }
}
