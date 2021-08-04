def partition(arr: list, lo: int, hi: int) -> None:
    if lo == hi:
        return

    pivot = arr[lo]

    left = lo
    right = hi - 1

    while left != right:
        for i in range(right, left, -1):
            if arr[i] < pivot:
                arr[left] = arr[i]
                arr[i] = pivot
                break
            right -= 1
        for i in range(left, right, +1):
            if arr[i] > pivot:
                arr[right] = arr[i]
                arr[i] = pivot
                break
            left += 1

    partition(arr, lo, left)
    partition(arr, left + 1, hi)

def sort(arr: list) -> None:
    partition(arr, 0, len(arr))
