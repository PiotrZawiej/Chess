package FS

type Rook struct {
	Mark            string
	HorizontalPlace int
	VerticalPlace   int
}

func NewRook(HP int, VP int) *Rook {
	r := new(Rook)
	r.Mark = "R "
	r.HorizontalPlace = HP
	r.VerticalPlace = VP

	return r
}