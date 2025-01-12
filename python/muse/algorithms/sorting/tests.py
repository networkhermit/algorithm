from collections.abc import Callable

from muse.util import sequence_builder, sequences


def derive(
    sort: Callable[[list], None], size: int, pack: Callable[[list], None]
) -> Callable[[], bool]:
    def f() -> bool:
        arr = [0] * size
        pack(arr)

        checksum = sequences.parity_checksum(arr)

        sort(arr)

        if sequences.parity_checksum(arr) != checksum:
            return False

        return sequences.is_sorted(arr)

    return f


def derive_decreasing(sort: Callable[[list], None], size: int) -> Callable[[], bool]:
    return derive(sort, size, sequence_builder.pack_decreasing)


def derive_empty(sort: Callable[[list], None]) -> Callable[[], bool]:
    return derive(sort, 0, sequence_builder.pack_identical)


def derive_identical(sort: Callable[[list], None], size: int) -> Callable[[], bool]:
    return derive(sort, size, sequence_builder.pack_identical)


def derive_increasing(sort: Callable[[list], None], size: int) -> Callable[[], bool]:
    return derive(sort, size, sequence_builder.pack_increasing)


def derive_random(sort: Callable[[list], None], size: int) -> Callable[[], bool]:
    return derive(sort, size, sequence_builder.pack_random)
