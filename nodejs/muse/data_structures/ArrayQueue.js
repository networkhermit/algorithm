'use strict'

exports.ArrayQueue = class {
  constructor (physicalSize = 0) {
    this.data = null
    this.front = 0
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
      throw new Error('[PANIC - NoSuchElement]')
    }

    return this.data[this.front]
  }

  enqueue (element) {
    if (this.logicalSize === this.physicalSize) {
      let newCapacity = 64

      if (this.physicalSize >= 64) {
        newCapacity = this.physicalSize + (this.physicalSize >>> 1)
      }

      const temp = new Array(newCapacity)

      let cursor = this.front

      for (let i = 0, length = this.logicalSize; i < length; i++) {
        if (cursor === this.physicalSize) {
          cursor = 0
        }
        temp[i] = this.data[cursor]
        cursor++
      }

      this.data = temp
      this.front = 0
      this.physicalSize = newCapacity
    }

    this.data[(this.front + this.logicalSize) % this.physicalSize] = element

    this.logicalSize++
  }

  dequeue () {
    if (this.logicalSize === 0) {
      throw new Error('[PANIC - NoSuchElement]')
    }

    this.data[this.front] = null

    this.front = (this.front + 1) % this.physicalSize

    this.logicalSize--
  }

  capacity () {
    return this.physicalSize
  }

  shrink () {
    const temp = new Array(this.logicalSize)

    let cursor = this.front

    for (let i = 0, length = this.logicalSize; i < length; i++) {
      if (cursor === this.physicalSize) {
        cursor = 0
      }
      temp[i] = this.data[cursor]
      cursor++
    }

    this.data = temp
    this.front = 0
    this.physicalSize = this.logicalSize
  }
}
