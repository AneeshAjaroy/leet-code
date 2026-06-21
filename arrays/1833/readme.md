# 1833. Maximum Ice Cream Bars

solving this problem is easy, just sort the costs, can count till aggregated sum is less then equal to the total coins preset

but we are asked to explicitly use the counting sort algorithm

so find the max cost ice cream, let it be i, create an array of size i+1,
and the populate it with the freq of each element, then create a new array with the data, this is the counting sort

and we can use this frwq array to get the count too