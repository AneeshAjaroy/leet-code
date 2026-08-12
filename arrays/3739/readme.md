# 3739. Count Subarrays With Majority Element II

In continuation with 3737, now the constraint is increased to 10^4 make the original soln with O(n^2) not feasable

if we build the prefic array, then we need to find all occurances such that prefix[i] < prefix[j], such that i<j

but the prefixes can be negative and random, if we can map them to natual numbers, then create a fenwick tree with freq, then the all occurances discussed above becomes the fenwick query