"use strict"

exports.ArrayStack = class {

    constructor (physicalSize = 0) {
        this.data = null
        this.logicalSize = 0
        this.physicalSize = 64

        if (physicalSize > 1) {
            this.physicalSize = physicalSize
        }
        this.data = new Array(this.physicalSize)
    }

    size () {
        return this.logicalSize
    }

    isEmpty () {
        return this.logicalSize === 0
    }

    peek () {
        if (this.logicalSize === 0) {
            throw new Error("[PANIC - NoSuchElement]")
        }

        return this.data[this.logicalSize - 1]
    }

    push (element) {
        if (this.logicalSize === this.physicalSize) {
            let newCapacity = 64

            if (this.physicalSize >= 64) {
                newCapacity = this.physicalSize + (this.physicalSize >>> 1)
            }

            const temp = new Array(newCapacity)

            for (let i = 0, length = this.logicalSize; i < length; i++) {
                temp[i] = this.data[i]
            }

            this.data = temp
            this.physicalSize = newCapacity
        }

        this.data[this.logicalSize] = element

        this.logicalSize++
    }

    pop () {
        if (this.logicalSize === 0) {
            throw new Error("[PANIC - NoSuchElement]")
        }

        this.logicalSize--

        this.data[this.logicalSize] = null
    }

    capacity () {
        return this.physicalSize
    }

    shrink () {
        const temp = new Array(this.logicalSize)

        for (let i = 0, length = this.logicalSize; i < length; i++) {
            temp[i] = this.data[i]
        }

        this.data = temp
        this.physicalSize = this.logicalSize
    }
}
