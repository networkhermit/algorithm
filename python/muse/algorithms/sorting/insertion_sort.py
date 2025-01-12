def sort(arr: list[int]) -> None:
    for i in range(1, len(arr)):
        target = arr[i]
        for cursor in range(i, 0, -1):
            if arr[cursor - 1] <= target:
                break
            arr[cursor] = arr[cursor - 1]
        else:
            cursor = 0
        arr[cursor] = target
