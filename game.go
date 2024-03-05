package main

import "fmt"

func gameRun() {
	board := boardStructure{}
		
	run := true
	for run {
		
		board.displayBoard()

		fmt.Println("Do you want to continue? (yes/no)")
		var input string
		fmt.Scanln(&input)

		if input != "yes"{
			run = false
		}

	}

	
}
