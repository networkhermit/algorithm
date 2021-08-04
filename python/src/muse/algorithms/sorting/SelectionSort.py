def sort(arr: list) -> None:
    if len(arr) == 0:
        return

    iMin = 0

    bound = len(arr) - 1
    for i in range(bound):
        iMin = i
        for j in range(i + 1, bound + 1):
            if arr[j] < arr[iMin]:
                iMin = j
        if iMin != i:
            arr[i], arr[iMin] = arr[iMin], arr[i]
