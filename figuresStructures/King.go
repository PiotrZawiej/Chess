package FS


type King struct{
	Mark string
	HorizontalPlace int
	VerticalPlace int
}

func NewKing(horizontalPlace int, verticalPLace int, mark string) *King{
	k := new(King)
	k.Mark = mark
	k.HorizontalPlace = horizontalPlace
	k.VerticalPlace = verticalPLace

	return k
}