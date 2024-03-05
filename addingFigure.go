package main

import "chess/chess/figuresStructures"


func (b *boardStructure) addPawn(p *FS.Pawn) {
    b.board[p.VerticalPlace][p.HorizontalPlace] = p.Mark
}

func (b *boardStructure) addAllPawns() {
    for i := 0; i < 8; i++ {
        pawn := FS.NewPawn(i, 1)
        b.addPawn(pawn)
    }

    for i := 0; i < 8; i++ {
        pawn := FS.NewPawn(i, 6)
        b.addPawn(pawn)
    }
}
