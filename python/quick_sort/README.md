The Quick Sort algorithm is a divide-and-conquer method for sorting. It works by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays, according to whether they are less than or greater than the pivot. The sub-arrays are then recursively sorted.

The steps are:

Choose an element from the array to serve as a pivot. Typically, this could be the first, middle, or last element.
Partition the other elements into two sub-arrays, according to whether they are less than or greater than the pivot.
Recursively apply the quicksort to the sub-arrays.
Unlike the binary search algorithm, the Quick Sort does not require the array to be sorted as it sorts the array itself. However, the choice of the pivot can greatly affect the performance. In the worst case (already sorted array and choosing the first or last element as pivot), it results in a time complexity of O(n^2). But in the best case (choosing the middle element as pivot on a 'random' array), it can achieve a time complexity of O(n log n).