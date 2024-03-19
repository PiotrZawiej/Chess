package FS

type Bishop struct{
	Mark string
	HorizontalPlace int
	VerticalPlace int
}


func NewBishop(VP int, HP int) *Bishop{
	b := new(Bishop)

	b.Mark = "B "
	b.HorizontalPlace = HP
	b.VerticalPlace = VP

	return b
}