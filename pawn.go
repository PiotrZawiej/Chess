package main

type pawn struct{
	mark rune
	horizontalPlace int8 
	verticalPlace int8
}

func NewPawn() *pawn{
	p := new(pawn)
	p.mark ='P'
	
	return p
}