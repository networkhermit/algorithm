"use strict"

const SinglyLinkedList = require("fun/vac/data_structures/SinglyLinkedList")
const TestRunner       = require("fun/vac/util/TestRunner")

const testSinglyLinkedList = () => {
    let size = 8192

    const list = new SinglyLinkedList.SinglyLinkedList()

    for (let i = 1; i <= size; i++) {
        list.append(i)
    }

    if (list.size() !== size) {
        return false
    }

    for (let i = 0; i < size; i++) {
        if (list.get(i) !== i + 1) {
            return false
        }
    }

    for (let i = 0; i < size; i++) {
        list.set(i, size - i)
    }

    for (let i = 0; i < size; i++) {
        if (list.get(i) !== size - i) {
            return false
        }
    }

    let x = 0, y = 0

    for (let i = 0, j = size - 1; i < j; i++, j--) {
        x = list.get(i)
        y = list.get(j)

        list.remove(i)
        list.insert(i, y)
        list.remove(j)
        list.insert(j, x)
    }

    for (let i = 0; i < size; i++) {
        if (list.get(i) !== i + 1) {
            return false
        }
    }

    for (let i = size; i >= 1; i--) {
        if (list.back() !== i) {
            return false
        } else {
            list.eject()
        }
    }

    if (!list.isEmpty()) {
        return false
    }

    return true
}

if (module === require.main) {
    TestRunner.parseTest(testSinglyLinkedList())
}
