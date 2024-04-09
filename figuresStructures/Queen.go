package FS


type Queen struct{
	Mark string
	HorizontalPlace int
	VerticalPlace int
}

func NewQueen(horizontalPlace int, verticalPLace int, mark string) *Queen{
	q := new(Queen)
	q.Mark = mark
	q.HorizontalPlace = horizontalPlace
	q.VerticalPlace = verticalPLace

	return q
}