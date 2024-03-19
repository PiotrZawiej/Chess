package FS

type Knight struct{
	Mark string
	HorizontalPlace int
	VerticalPlace int
}

func NewKnight(VP int, HP int) *Knight{
	k := new(Knight)

	k.Mark = "K "
	k.HorizontalPlace = HP
	k.VerticalPlace = VP

	return k
}

