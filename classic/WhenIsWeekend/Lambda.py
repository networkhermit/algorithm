import time


def main() -> None:
    print("When’s Weekend?")

    today = time.localtime().tm_wday

    match today:
        case 6:
            print("Today.")
        case 5:
            print("Today.")
        case 4:
            print("Tomorrow.")
        case 3:
            print("In two days.")
        case _:
            print("Too far away.")


if __name__ == "__main__":
    main()
