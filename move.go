package main

import "fmt"

type move struct {
	boardNumber int
	boardLetter byte
}

func chooseFigure() *move {
	fmt.Println("Select figure")
	fmt.Scanln()
}
