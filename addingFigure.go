package main

import FS "chess/chess/figuresStructures"

func (b *boardStructure) addPawn(p *FS.Pawn) {
	b.board[p.VerticalPlace][p.HorizontalPlace] = p.Mark
}

func (b *boardStructure) addAllPawns() {
	for i := 0; i < 8; i++ {
		pawn := FS.NewPawn(i, 1, "♟ ")
		b.addPawn(pawn)
	}

	for i := 0; i < 8; i++ {
		pawn := FS.NewPawn(i, 6, "♙ ")
		b.addPawn(pawn)
	}
}

func (b *boardStructure) addRook(r *FS.Rook) {
	b.board[r.HorizontalPlace][r.VerticalPlace] = r.Mark
}

func (b *boardStructure) addAllRooks() {
	rook1 := FS.NewRook(0, 0, "♜ ")
	b.addRook(rook1)

	rook2 := FS.NewRook(0, len(b.board)-1, "♜ ")
	b.addRook(rook2)

	rook3 := FS.NewRook(len(b.board)-1, 0, "♖ ")
	b.addRook(rook3)

	rook4 := FS.NewRook(len(b.board)-1, len(b.board)-1, "♖ ")
	b.addRook(rook4)

}

func (b *boardStructure) addKnight(k *FS.Knight) {
	b.board[k.HorizontalPlace][k.VerticalPlace] = k.Mark
}

func (b *boardStructure) addAllKnights() {
	knight1 := FS.NewKnight(1, 0, "♞ ")
	b.addKnight(knight1)

	knight2 := FS.NewKnight(1, len(b.board)-1, "♘ ")
	b.addKnight(knight2)

	knight3 := FS.NewKnight(len(b.board)-2, 0, "♞ ")
	b.addKnight(knight3)

	knight4 := FS.NewKnight(len(b.board)-2, len(b.board)-1, "♘ ")
	b.addKnight(knight4)
}

func (b *boardStructure) addBishop(bi *FS.Bishop) {
	b.board[bi.HorizontalPlace][bi.VerticalPlace] = bi.Mark
}

func (b *boardStructure) addAllBishops() {
	bishop1 := FS.NewBishop(2, 0, "♝ ")
	b.addBishop(bishop1)

	bishop2 := FS.NewBishop(2, len(b.board)-1, "♗ ")
	b.addBishop(bishop2)

	bishop3 := FS.NewBishop(len(b.board)-3, 0, "♝ ")
	b.addBishop(bishop3)

	bishop4 := FS.NewBishop(len(b.board)-3, len(b.board)-1, "♗ ")
	b.addBishop(bishop4)

}

func (b *boardStructure) addQueen(q *FS.Queen) {
	b.board[q.HorizontalPlace][q.VerticalPlace] = q.Mark
}

func (b *boardStructure) addAllQueens() {
	BlackQueen := FS.NewQueen(0, 3, "♛ ")
	b.addQueen(BlackQueen)

	WhiteQueen := FS.NewQueen(7, 3, "♕ ")
	b.addQueen(WhiteQueen)
}

func (b *boardStructure) addKing(k *FS.King) {
	b.board[k.HorizontalPlace][k.VerticalPlace] = k.Mark
}

func (b *boardStructure) addAllKings() {
	BlackKing := FS.NewKing(0, 4, "♚ ")
	b.addKing(BlackKing)

	WhiteKing := FS.NewKing(7, 4, "♔ ")
	b.addKing(WhiteKing)
}
