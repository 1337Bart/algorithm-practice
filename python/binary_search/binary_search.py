import random
import time


def generate_sorted_list(n):
    lst = [random.randint(0, n) for _ in range(n)]
    lst.sort()  # Ensure list is sorted
    return lst


def binary_search(haystack, needle):
    low = 0
    high = len(haystack) - 1

    while low <= high:
        mid = (low + high) // 2
        if haystack[mid] < needle:
            low = mid + 1
        elif haystack[mid] > needle:
            high = mid - 1
        else:
            return True  # needle found
    return False  # needle not found


def main():
    haystack = generate_sorted_list(1000000)
    needle = random.choice(haystack)  # Ensure the needle is in the haystack

    start_time = time.time()
    print(binary_search(haystack, needle))
    duration = time.time() - start_time

    print(f"Execution time: {duration*1000} ms")


if __name__ == "__main__":
    main()
