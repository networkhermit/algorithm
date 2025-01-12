from collections.abc import Callable

from muse.algorithms.number_theory import factorial
from muse.util import test_runner

sample = [
    [0, 1],
    [1, 1],
    [2, 2],
    [3, 6],
    [4, 24],
    [5, 120],
    [6, 720],
    [7, 5040],
    [8, 40320],
    [9, 362880],
    [10, 3_628_800],
    [11, 39_916_800],
    [12, 479_001_600],
]


def derive(fn: Callable[[int], int]) -> Callable[[], bool]:
    def f() -> bool:
        for tc in sample:
            if fn(tc[0]) != tc[1]:
                return False
        return True

    return f


def test_factorial() -> bool:
    return (
        derive(factorial.iterative_procedure)()
        and derive(factorial.recursive_procedure)()
    )


def main() -> None:
    test_runner.pick(test_factorial)


if __name__ == "__main__":
    main()
