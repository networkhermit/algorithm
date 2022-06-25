from muse.algorithms.number_theory import factorial
from muse.util import test_runner


def test_factorial() -> bool:
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

    for i, _ in enumerate(sample):
        if factorial.iterative_procedure(sample[i][0]) != sample[i][1]:
            return False

    for i, _ in enumerate(sample):
        if factorial.recursive_procedure(sample[i][0]) != sample[i][1]:
            return False

    return True


def main() -> None:
    test_runner.pick(test_factorial)


if __name__ == "__main__":
    main()
