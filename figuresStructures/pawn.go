package FS


type Pawn struct{
	Mark string
	HorizontalPlace int
	VerticalPlace int
}

func NewPawn(horizontalPlace int, verticalPLace int, mark string) *Pawn{
	p := new(Pawn)
	p.Mark = mark
	p.HorizontalPlace = horizontalPlace
	p.VerticalPlace = verticalPLace

	return p
}
