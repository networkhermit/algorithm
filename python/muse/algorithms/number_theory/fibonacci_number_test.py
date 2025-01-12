from collections.abc import Callable

from muse.algorithms.number_theory import fibonacci_number
from muse.util import test_runner

sample = [
    [-31, 1_346_269],
    [-30, -832040],
    [-29, 514229],
    [-28, -317811],
    [-27, 196418],
    [-26, -121393],
    [-25, 75025],
    [-24, -46368],
    [-23, 28657],
    [-22, -17711],
    [-21, 10946],
    [-20, -6765],
    [-19, 4181],
    [-18, -2584],
    [-17, 1597],
    [-16, -987],
    [-15, 610],
    [-14, -377],
    [-13, 233],
    [-12, -144],
    [-11, 89],
    [-10, -55],
    [-9, 34],
    [-8, -21],
    [-7, 13],
    [-6, -8],
    [-5, 5],
    [-4, -3],
    [-3, 2],
    [-2, -1],
    [-1, 1],
    [0, 0],
    [1, 1],
    [2, 1],
    [3, 2],
    [4, 3],
    [5, 5],
    [6, 8],
    [7, 13],
    [8, 21],
    [9, 34],
    [10, 55],
    [11, 89],
    [12, 144],
    [13, 233],
    [14, 377],
    [15, 610],
    [16, 987],
    [17, 1597],
    [18, 2584],
    [19, 4181],
    [20, 6765],
    [21, 10946],
    [22, 17711],
    [23, 28657],
    [24, 46368],
    [25, 75025],
    [26, 121393],
    [27, 196418],
    [28, 317811],
    [29, 514229],
    [30, 832040],
    [31, 1_346_269],
]


def derive(fn: Callable[[int], int]) -> Callable[[], bool]:
    def f() -> bool:
        for tc in sample:
            if fn(tc[0]) != tc[1]:
                return False
        return True

    return f


def test_fibonacci_number() -> bool:
    return (
        derive(fibonacci_number.iterative_procedure)()
        and derive(fibonacci_number.recursive_procedure)()
    )


def main() -> None:
    test_runner.pick(test_fibonacci_number)


if __name__ == "__main__":
    main()
