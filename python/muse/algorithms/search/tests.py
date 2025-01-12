from collections.abc import Callable

from muse.util import sequence_builder


def derive(
    find: Callable[[list, object], int], size: int, pack: Callable[[list], None]
) -> Callable[[], bool]:
    def f() -> bool:
        arr = [0] * size
        pack(arr)

        sentinel = [-1, 2_147_483_647]

        for v in sentinel:
            if find(arr, v) != size:
                return False

        for i, v in enumerate(arr):
            if find(arr, v) != i:
                return False

        return True

    return f


def derive_empty(find: Callable[[list, object], int]) -> Callable[[], bool]:
    return derive(find, 0, sequence_builder.pack_identical)


def derive_increasing(
    find: Callable[[list, object], int], size: int
) -> Callable[[], bool]:
    return derive(find, size, sequence_builder.pack_increasing)
