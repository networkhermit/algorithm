def fizz_buzz(n: int) -> str:
    word = ""

    if n % 3 == 0:
        word += "Fizz"
    if n % 5 == 0:
        word += "Buzz"

    if len(word) == 0:
        word = str(n)

    return word


def main() -> None:
    for i in range(1, 101):
        print(fizz_buzz(i))


if __name__ == "__main__":
    main()
