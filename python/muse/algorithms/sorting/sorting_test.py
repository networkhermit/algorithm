from muse.algorithms import sorting
from muse.util import sequence_builder, sequences, test_runner


def test_bubble_sort() -> bool:
    size = 32768

    arr = [0] * size
    sequence_builder.pack_random(arr)

    checksum = sequences.parity_checksum(arr)

    sorting.bubblesort(arr)

    if sequences.parity_checksum(arr) != checksum:
        return False

    return sequences.is_sorted(arr)


def test_insertion_sort() -> bool:
    size = 32768

    arr = [0] * size
    sequence_builder.pack_random(arr)

    checksum = sequences.parity_checksum(arr)

    sorting.insertionsort(arr)

    if sequences.parity_checksum(arr) != checksum:
        return False

    return sequences.is_sorted(arr)


def test_merge_sort() -> bool:
    size = 32768

    arr = [0] * size
    sequence_builder.pack_random(arr)

    checksum = sequences.parity_checksum(arr)

    sorting.mergesort(arr)

    if sequences.parity_checksum(arr) != checksum:
        return False

    return sequences.is_sorted(arr)


def test_quick_sort() -> bool:
    size = 32768

    arr = [0] * size
    sequence_builder.pack_random(arr)

    checksum = sequences.parity_checksum(arr)

    sorting.quicksort(arr)

    if sequences.parity_checksum(arr) != checksum:
        return False

    return sequences.is_sorted(arr)


def test_selection_sort() -> bool:
    size = 32768

    arr = [0] * size
    sequence_builder.pack_random(arr)

    checksum = sequences.parity_checksum(arr)

    sorting.selectionsort(arr)

    if sequences.parity_checksum(arr) != checksum:
        return False

    return sequences.is_sorted(arr)


def main() -> None:
    test_runner.pick(test_bubble_sort)

    test_runner.pick(test_insertion_sort)

    test_runner.pick(test_merge_sort)

    test_runner.pick(test_quick_sort)

    test_runner.pick(test_selection_sort)


if __name__ == "__main__":
    main()
