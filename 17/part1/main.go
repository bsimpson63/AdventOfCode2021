package main

/*
The probe's x,y position starts at 0,0. Then, it will follow some trajectory by moving in steps. On each step, these changes occur in the following order:

The probe's x position increases by its x velocity.
The probe's y position increases by its y velocity.
Due to drag, the probe's x velocity changes by 1 toward the value 0;
  that is, it decreases by 1 if it is greater than 0, increases by 1
  if it is less than 0, or does not change if it is already 0.
Due to gravity, the probe's y velocity decreases by 1.

Find the initial velocity that causes the probe to reach the highest
y position and still eventually be within the target area after any step.
What is the highest y position it reaches on this trajectory?

y and x are independent

can we calculate vy first?
say vy is 10:
0: 0,vy:10
1: y:10,vy:9
2: y:19,vy:8
3: ..

we always go back through y=0, and at that point our vy=-(vy0+1)
so that means our next location will be y=-vy0-1


*/

func main() {
	// simple example
	// target area: x=20..30, y=-10..-5
	//targetX := [2]int{20, 20}
	//targetY := [2]int{-10, -5}
	// real input
	// target area: x=244..303, y=-91..-54
	//targetX := [2]int{244, 303}
	//targetY := [2]int{-91, -54}
	// if our target is y=-91 that means we can go vy0=90
	// using a spreadsheet we can see that ymax is 4095
}
