from muse.algorithms import search
from muse.algorithms.search import tests
from muse.util import test_runner


def test_binary_search() -> bool:
    return tests.derive_increasing(search.binarysearch, 100000)()


def test_linear_search() -> bool:
    return tests.derive_increasing(search.linearsearch, 10000)()


def main() -> None:
    test_runner.pick(test_binary_search)

    test_runner.pick(test_linear_search)


if __name__ == "__main__":
    main()
