package FS

type Knight struct{
	Mark string
	HorizontalPlace int
	VerticalPlace int
}

func NewKnight(VP int, HP int, mark string) *Knight{
	k := new(Knight)

	k.Mark = mark
	k.HorizontalPlace = HP
	k.VerticalPlace = VP

	return k
}

