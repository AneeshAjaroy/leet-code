# 1840. Maximum Building Height

okay, to we have to find the maximum possible height 

I think we can use greedy approach
set height of the first 2 buildings 0,1 (restriction)
then set the maximum height of buildings according to restriction points

now start from the first building, see it it is already set, if yes, and satisfies, proceed foreward
if not satisfies, then reduce it to height og prev buildign + 1 and proceed if the height is larger than prev
if height is lower than prev, the set the height of prev building as height of current + 1 and then again move backward to previous building, to check its correctness



then give out the maximum height

The soln is correct, but out 10^9, as the height array is too large

optimized soln

we can see that restrictions mutually affetc each other
(i,hi),(j,hj),(k,hk)
height at j due to i is min(h1+j-1,hj)
similarly height at j due to k is min(kh+k-j,hj)

so we can move foreward/backward to get a completely satisfing restriction set (ofc,i thnk we must first sort them)
once we have it, we cna find the height of largest accomadable building in between as hj+hi+j-i/2

as there is sorting time complexity is o(nlogn)