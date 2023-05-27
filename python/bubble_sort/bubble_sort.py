def bubble_sort(haystack):
    n = len(haystack)
    for i in range(n):
        for j in range(0, n - i - 1):
            if haystack[j] > haystack[j + 1]:
                haystack[j], haystack[j + 1] = haystack[j + 1], haystack[j]
