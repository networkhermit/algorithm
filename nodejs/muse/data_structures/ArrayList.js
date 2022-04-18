'use strict'

exports.ArrayList = class {
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

  get (index) {
    if (index < 0 || index >= this.logicalSize) {
      throw new Error('[PANIC - IndexOutOfBounds]')
    }

    return this.data[index]
  }

  set (index, element) {
    if (index < 0 || index >= this.logicalSize) {
      throw new Error('[PANIC - IndexOutOfBounds]')
    }

    this.data[index] = element
  }

  insert (index, element) {
    if (index < 0 || index > this.logicalSize) {
      throw new Error('[PANIC - IndexOutOfBounds]')
    }

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

    for (let i = this.logicalSize; i > index; i--) {
      this.data[i] = this.data[i - 1]
    }

    this.data[index] = element

    this.logicalSize++
  }

  remove (index) {
    if (index < 0 || index >= this.logicalSize) {
      throw new Error('[PANIC - IndexOutOfBounds]')
    }

    for (let i = index + 1; i < this.logicalSize; i++) {
      this.data[i - 1] = this.data[i]
    }

    this.logicalSize--

    this.data[this.logicalSize] = null
  }

  front () {
    return this.get(0)
  }

  back () {
    return this.get(this.logicalSize - 1)
  }

  prepend (element) {
    this.insert(0, element)
  }

  append (element) {
    this.insert(this.logicalSize, element)
  }

  poll () {
    this.remove(0)
  }

  eject () {
    this.remove(this.logicalSize - 1)
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
