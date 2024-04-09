package FS

type Bishop struct{
	Mark string
	HorizontalPlace int
	VerticalPlace int
}


func NewBishop(VP int, HP int, mark  string) *Bishop{
	b := new(Bishop)

	b.Mark = mark
	b.HorizontalPlace = HP
	b.VerticalPlace = VP

	return b
}