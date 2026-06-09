# 3689. Maximum Total Subarray Value I

looking at the problem, we can choose subarray of any length, but must choose k subarrays

since we want max sum value, we want to maximise each individual value
max of a individual value can be max-min of array,

since subarrays can be same => result == k*(max num -min num)

since we loop over once to find max and min
Time Complexity => O(n)