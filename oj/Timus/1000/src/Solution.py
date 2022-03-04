import sys

def main() -> None:
    for line in sys.stdin:
        pair = line.split()
        a = int(pair[0])
        b = int(pair[1])
        print(a + b)

if __name__ == "__main__":
    main()
