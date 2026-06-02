# 3633. Earliest Finish Time for Land and Water Rides I

I think the key is solving both cases water ride + land ride / land ride + water ride
and select the best one

solving water ride + land ride 

add water ride and completion times and select the one with lowest value, which is water ride end time

next loop over land rides and if the ride is already started and its duration or wait for it to start and then end and select the min time for this

this is the one answer value

repeat the same with land ride + water ride

select the least of two


normal looping over arrays only

Time Complexity :- O(n)
