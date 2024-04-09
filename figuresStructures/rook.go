package FS

type Rook struct {
	Mark            string
	HorizontalPlace int
	VerticalPlace   int
}

func NewRook(HP int, VP int, mark string) *Rook {
	r := new(Rook)
	r.Mark = mark
	r.HorizontalPlace = HP
	r.VerticalPlace = VP

	return r
}