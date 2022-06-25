from muse.algorithms.sorting import selection_sort
from muse.util import sequence_builder, sequences, test_runner


def test_selection_sort() -> bool:
    size = 32768

    arr = [0] * size
    sequence_builder.pack_random(arr)

    checksum = sequences.parity_checksum(arr)

    selection_sort.sort(arr)

    if sequences.parity_checksum(arr) != checksum:
        return False

    return sequences.is_sorted(arr)


def main() -> None:
    test_runner.pick(test_selection_sort)


if __name__ == "__main__":
    main()
