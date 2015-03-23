package squareone

// A Color is a number in the range [0, 6).
type Color int

// A Piece is an edge or corner piece.
type Piece struct {
	Edge        bool
	TopColor    int
	SecondColor int
	ThirdColor  int
}

// A Square1 represents the state of a square-1 puzzle. It does not allow for
// situations where a 60 degree piece is halfway between the middle crack.
type Square1 struct {
	MiddleSquare bool
	BottomCount  int
	BottomPieces [10]Piece
	TopCount     int
	TopPieces    [10]Piece
}

// TurnBottom turns the bottom the next reasonable amount.
// This returns 1 or 2 depending on whether it moved an edge or a corner past
// the crack, respectively.
func (s Square1) TurnBottom() int {
	piece := s.BottomPieces[s.BottomCount - 1]
	copy(s.BottomPieces[1:], s.BottomPieces[0:])
	s.BottomPieces[0] = piece
	if piece.Edge {
		return 1
	} else {
		return 2
	}
}

// TurnTop turns the top the next reasonable amount.
// This returns 1 or 2 depending on whether it moved an edge or a corner past
// the crack, respectively.
func (s Square1) TurnTop() int {
	piece := s.TopPieces[s.TopCount - 1]
	copy(s.TopPieces[1:], s.TopPieces[0:])
	s.TopPieces[0] = piece
	if piece.Edge {
		return 1
	} else {
		return 2
	}
}
