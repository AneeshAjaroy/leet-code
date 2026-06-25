# 3737. Count Subarrays With Majority Element I

okay, we can convert this problem such that it can become a count number of subsets with sum grater than 0

we can replace every target with q and every non target with -1, as the requirement is target freq strictly greater than half, we can count subarrays of new array with sum greater than 0

with constraint n < 1000, after building the new array we can loop through all subarrays => O(n^2)