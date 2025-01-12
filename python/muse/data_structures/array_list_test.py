from muse.data_structures import tests
from muse.data_structures.array_list import ArrayList
from muse.util import test_runner


def test_array_list() -> bool:
    return tests.list_derive(ArrayList)()


def main() -> None:
    test_runner.pick(test_array_list)


if __name__ == "__main__":
    main()
