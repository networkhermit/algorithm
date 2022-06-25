from muse.algorithms.sorting import insertion_sort
from muse.util import sequence_builder, sequences, test_runner


def test_insertion_sort() -> bool:
    size = 32768

    arr = [0] * size
    sequence_builder.pack_random(arr)

    checksum = sequences.parity_checksum(arr)

    insertion_sort.sort(arr)

    if sequences.parity_checksum(arr) != checksum:
        return False

    return sequences.is_sorted(arr)


def main() -> None:
    test_runner.pick(test_insertion_sort)


if __name__ == "__main__":
    main()
