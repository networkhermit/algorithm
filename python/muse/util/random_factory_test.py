from muse.util import random_factory, test_runner


def test_gen_int_n() -> bool:
    value = 0
    for _ in range(8192):
        if random_factory.gen_int_n(0, 0) != 0:
            return False

        if random_factory.gen_int_n(1, 1) != 1:
            return False

        value = random_factory.gen_int_n(0, 1)
        if value < 0 or value > 1:
            return False

        value = random_factory.gen_int_n(100, 10000)
        if random_factory.gen_int_n(value, value) != value:
            return False
        if value < 100 or value > 10000:
            return False

    return True


def test_gen_even() -> bool:
    for _ in range(8192):
        if (random_factory.gen_even() & 1) != 0:
            return False

    return True


def test_gen_odd() -> bool:
    for _ in range(8192):
        if (random_factory.gen_odd() & 1) == 0:
            return False

    return True


def main() -> None:
    test_runner.pick(test_gen_int_n)

    test_runner.pick(test_gen_even)

    test_runner.pick(test_gen_odd)


if __name__ == "__main__":
    main()
