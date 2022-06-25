from muse.util import test_runner


def test_pick() -> None:
    for i in range(10):
        test_runner.pick(lambda: (i & 1) != 0)


def main() -> None:
    test_pick()


if __name__ == "__main__":
    main()
