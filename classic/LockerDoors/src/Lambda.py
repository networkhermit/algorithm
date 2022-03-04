import math

def bruteForce(n: int) -> list:
    arr = [False] * n

    for i in range(1, n + 1):
        for j in range(i - 1, n, i):
            arr[j] = not arr[j]

    return arr

def process(n: int) -> list:
    arr = [False] * n

    temp = 0
    for i in range(1, n + 1):
        temp = int(math.sqrt(float(i)))
        arr[i - 1] = temp * temp == i

    return arr

def main() -> None:
    N = 100

    arr = process(N)

    for i, v in enumerate(arr):
        if v:
            print("%d\t  OPEN" % (i + 1))

if __name__ == "__main__":
    main()
