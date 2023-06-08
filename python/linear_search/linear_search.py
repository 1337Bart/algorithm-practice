import random
import time


def generate_list(n):
    lst = [random.randint(0, n) for _ in range(n)]
    return lst


def linear_search(haystack, needle):
    for i in range(len(haystack)):
        if haystack[i] == needle:
            print("Needle found")


def main():
    haystack = generate_list(1000)
    needle = random.choice(haystack)  # Ensure the needle is in the haystack

    start_time = time.time()
    linear_search(haystack, needle)
    duration = time.time() - start_time

    print(f"Execution time: {duration*1000} ms")


if __name__ == "__main__":
    main()
