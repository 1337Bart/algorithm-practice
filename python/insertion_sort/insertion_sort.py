def insertion_sort(haystack):
    for i in range(1, len(haystack)):
        needle = haystack[i]
        j = i - 1
        while j >= 0 and needle < haystack[j]:
            haystack[j + 1] = haystack[j]
            j -= 1
        haystack[j + 1] = needle
