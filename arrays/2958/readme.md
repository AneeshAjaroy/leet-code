# 2958

So the problem is to find the length of the longest subarray where the freq of each lement is less than equal to k

this seems to be similar to a dynamic sliding window and we need to find the length of the longest valid window

we will have start and end of window

we extend the end towards right as long as the invariant statiesfies 

if at any place the invariant is broken we move the start till the invariant destroyer value is crossed