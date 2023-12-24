var('x y z vx vy vz t1 t2 t3 ans')
eq1 = x + (vx * t1) == 19 + (-2 * t1)
eq2 = y + (vy * t1) == 13 + (1 * t1)
eq3 = z + (vz * t1) == 30 + (-2 * t1)
eq4 = x + (vx * t2) == 18 + (-1 * t2)
eq5 = y + (vy * t2) == 19 + (-1 * t2)
eq6 = z + (vz * t2) == 22 + (-2 * t2)
eq7 = x + (vx * t3) == 20 + (-2 * t3)
eq8 = y + (vy * t3) == 25 + (-2 * t3)
eq9 = z + (vz * t3) == 34 + (-4 * t3)
eq10 = ans == x + y + z
print(solve([eq1,eq2,eq3,eq4,eq5,eq6,eq7,eq8,eq9,eq10],x,y,z,vx,vy,vz,t1,t2,t3,ans))

