The binary search algorithm works by repeatedly dividing the search space in half. It starts by comparing the target value to the middle element of the array. If the target value is equal to the middle element, then the position is returned. If the target value is less than the middle element, the search continues on the lower half of the array. If the target value is greater than the middle element, the search continues on the upper half.

This halving of the search space continues until the value is found or the search space is empty.
This process assumes that the array is sorted. If it wasn't, there would be no guarantee that the value less than or greater than the middle element should be in the lower or upper half of the array, respectively.

So if you are using binary search, the list (or array) must be sorted for the algorithm to work correctly.


The time complexity of the algorithm is O(log n), where n is the size of the list.