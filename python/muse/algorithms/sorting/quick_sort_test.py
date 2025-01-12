from muse.algorithms.sorting import quick_sort, tests
from muse.util import test_runner


def test_quick_sort() -> bool:
    sort = quick_sort.sort
    return (
        tests.derive_decreasing(sort, 100)()
        and tests.derive_empty(sort)()
        and tests.derive_identical(sort, 100)()
        and tests.derive_increasing(sort, 100)()
        and tests.derive_random(sort, 100000)()
    )


def main() -> None:
    test_runner.pick(test_quick_sort)


if __name__ == "__main__":
    main()
