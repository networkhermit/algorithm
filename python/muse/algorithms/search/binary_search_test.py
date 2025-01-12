from muse.algorithms.search import binary_search, tests
from muse.util import test_runner


def test_binary_search() -> bool:
    find = binary_search.find
    return tests.derive_empty(find)() and tests.derive_increasing(find, 100000)()


def main() -> None:
    test_runner.pick(test_binary_search)


if __name__ == "__main__":
    main()
