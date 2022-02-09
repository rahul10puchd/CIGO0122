This is to demonstrate how fast GOROUTINES can make a program execute.

There are three prgrams written in GoLang -
	1. main.go (no goroutines are used in this)
	2. con_main.go (concurrency is used in this but only to have two seperate goroutines of the function, it doea make the program execution a bit faster but it is still slow)
	3. super_con_main.go (concurrency is not only used in calling the function but also used in the loops. Due to this we are performing the loop task by creating a seperate goroutine for each loop. By doing this the program execution becomes much faster compare to the other implementations done in above GO program)

	
By running all the programs we can see the time difference between the execution and can clearly see and conclude how fast GoROutines exactly are.
