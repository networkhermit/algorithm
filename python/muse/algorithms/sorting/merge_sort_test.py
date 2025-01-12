from muse.algorithms.sorting import merge_sort, tests
from muse.util import test_runner


def test_merge_sort() -> bool:
    sort = merge_sort.sort
    return (
        tests.derive_decreasing(sort, 100000)()
        and tests.derive_empty(sort)()
        and tests.derive_identical(sort, 100000)()
        and tests.derive_increasing(sort, 100000)()
        and tests.derive_random(sort, 100000)()
    )


def main() -> None:
    test_runner.pick(test_merge_sort)


if __name__ == "__main__":
    main()
