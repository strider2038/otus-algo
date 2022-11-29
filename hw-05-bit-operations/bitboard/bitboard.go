package bitboard

const (
	excludeA = uint64(0xFEFEFEFEFEFEFEFE)
	excludeB = uint64(0xFDFDFDFDFDFDFDFD)
	excludeC = uint64(0xFBFBFBFBFBFBFBFB)
	excludeD = uint64(0xF7F7F7F7F7F7F7F7)
	excludeE = uint64(0xEFEFEFEFEFEFEFEF)
	excludeF = uint64(0xDFDFDFDFDFDFDFDF)
	excludeG = uint64(0xBFBFBFBFBFBFBFBF)
	excludeH = uint64(0x7F7F7F7F7F7F7F7F)
)

var masks = map[int]uint64{
	-7: excludeA & excludeB & excludeC & excludeD & excludeE & excludeF & excludeG,
	-6: excludeA & excludeB & excludeC & excludeD & excludeE & excludeF,
	-5: excludeA & excludeB & excludeC & excludeD & excludeE,
	-4: excludeA & excludeB & excludeC & excludeD,
	-3: excludeA & excludeB & excludeC,
	-2: excludeA & excludeB,
	-1: excludeA,
	1:  excludeH,
	2:  excludeH & excludeG,
	3:  excludeH & excludeG & excludeF,
	4:  excludeH & excludeG & excludeF & excludeE,
	5:  excludeH & excludeG & excludeF & excludeE & excludeD,
	6:  excludeH & excludeG & excludeF & excludeE & excludeD & excludeC,
	7:  excludeH & excludeG & excludeF & excludeE & excludeD & excludeC & excludeB,
}

func KingMoves(index uint8) (int, uint64) {
	p := uint64(1) << index
	pa := p & excludeA
	ph := p & excludeH
	moves := pa<<7 | p<<8 | ph<<9 |
		/**/ pa>>1 | /*  */ ph<<1 |
		/**/ pa>>9 | p>>8 | ph>>7

	return CountBitsByDivision(moves), moves
}

func KnightMoves(index uint8) (int, uint64) {
	p := uint64(1) << index
	pa := p & excludeA
	pab := p & excludeA & excludeB
	pgh := p & excludeG & excludeH
	ph := p & excludeH
	moves := pa<<15 | ph<<17 |
		pab<<6 | pgh<<10 |
		pab>>10 | pgh>>6 |
		pa>>17 | ph>>15

	return CountBitsByDivision(moves), moves
}

func RookMoves(index uint8) (int, uint64) {
	p := uint64(1) << index
	moves := uint64(0)

	for i := uint64(1); i <= 8; i++ {
		moves |= p<<(8*i) | // ход вверх
			p>>(8*i) | // ход вниз
			(p&masks[int(-i)])>>i | // ход влево
			(p&masks[int(i)])<<i // ход вправо
	}

	return CountBitsByDivision(moves), moves
}
