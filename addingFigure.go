package main


func (b *boardStructure) addPawn(p *pawn) {
    b.board[p.verticalPlace][p.horizontalPlace] = string(p.mark)
}

func (b *boardStructure) addAllPawns(){
	for i := 0; i < 8; i++{
		pawn := NewPawn(i,1)
		b.addPawn(pawn)
	}

	for i := 0; i < 8; i++{
		pawn := NewPawn(i,6)
		b.addPawn(pawn)
	}
}