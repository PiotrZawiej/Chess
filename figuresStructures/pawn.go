package FS


type Pawn struct{
	Mark string
	HorizontalPlace int
	VerticalPlace int
}

func NewPawn(horizontalPlace int, verticalPLace int) *Pawn{
	p := new(Pawn)
	p.Mark ="P "
	p.HorizontalPlace = horizontalPlace
	p.VerticalPlace = verticalPLace

	return p
}
