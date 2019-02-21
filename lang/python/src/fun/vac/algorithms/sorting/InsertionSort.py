def sort(arr: list) -> None:
    target = None

    cursor = 0

    for i in range(1, len(arr)):
        target = arr[i]
        cursor = i
        while cursor > 0:
            if arr[cursor - 1] > target:
                arr[cursor] = arr[cursor - 1]
            else:
                break
            cursor -= 1
        arr[cursor] = target
