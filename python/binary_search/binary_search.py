import random
import time


def generate_sorted_list(n):
    lst = [random.randint(0, n) for _ in range(n)]
    lst.sort()  # Ensure list is sorted
    return lst


def binary_search(sorted_list, target_value):
    lower_bound = 0
    upper_bound = len(sorted_list) - 1

    # Continue while there are still elements to search within the bounds.
    while lower_bound <= upper_bound:
        middle_index = (lower_bound + upper_bound) // 2

        # If the middle value is less than the target, discard the lower half of the list.
        if sorted_list[middle_index] < target_value:
            lower_bound = middle_index + 1
        # If the middle value is more than the target, discard the upper half of the list.
        elif sorted_list[middle_index] > target_value:
            upper_bound = middle_index - 1
        else:
            return True  # Target value found

    return False  # Target value not found


def main():
    haystack = generate_sorted_list(1000000)
    needle = random.choice(haystack)  # Ensure the needle is in the haystack

    start_time = time.time()
    found = binary_search(haystack, needle)
    if found:
        print("Needle found.")
    duration = time.time() - start_time

    print(f"Execution time: {duration*1000} ms")
