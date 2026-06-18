# 1344. Angle Between Hands of a Clock

so the problem is very simple

hour hand 360 deg in 12 * 60 min => 0.5 deg in 1 min
minute hand 360 deg in 60 min => 6 deg in 1 min

take the line pointing towards 12 as the base 
calculate total rotation of hour and total rotation of min
then take diff
if its negative, the absolute value of it
if its greater then 180, then take the other angle obtained by sub from 360