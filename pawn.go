package main

type pawn struct{
	mark string
	horizontalPlace int
	verticalPlace int
}

func NewPawn(horizontalPlace int, verticalPLace int) *pawn{
	p := new(pawn)
	p.mark ="P "
	p.horizontalPlace = horizontalPlace
	p.verticalPlace = verticalPLace

	return p
}
