from muse.algorithms.search import linear_search, tests
from muse.util import test_runner


def test_linear_search() -> bool:
    find = linear_search.find
    return tests.derive_empty(find)() and tests.derive_increasing(find, 10000)()


def main() -> None:
    test_runner.pick(test_linear_search)


if __name__ == "__main__":
    main()
