import random
import time


def quick_sort(lst):
    # Base case: if the list contains 0 or 1 element, it is already sorted.
    if len(lst) <= 1:
        return lst
    else:
        # Choose a pivot randomly and remove it from the list.
        pivot = lst.pop(random.randint(0, len(lst) - 1))

        # Divide the remaining elements into two lists based on whether they are smaller or larger than the pivot.
        smaller = [i for i in lst if i < pivot]
        larger = [i for i in lst if i >= pivot]

        # Recursively sort the smaller and larger lists and combine the results.
        return quick_sort(smaller) + [pivot] + quick_sort(larger)


def main():
    # Generate a random list to sort.
    unsorted_list = [random.randint(0, 1000) for _ in range(1000)]
    print(f"Unsorted list: {unsorted_list}")

    # Sort the list.
    start_time = time.time()
    sorted_list = quick_sort(unsorted_list)
    duration = time.time() - start_time

    print(f"Execution time: {duration*1000} ms")
    print(f"Sorted list: {sorted_list}")


if __name__ == "__main__":
    main()
