package squareone

// A Color is a number in the range [0, 6).
type Color int

const (
	Yellow = 0
	White  = 1
	Blue   = 2
	Green  = 3
	Red    = 5
	Orange = 6
)

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
	Bottom       [10]Piece
	TopCount     int
	Top          [10]Piece
}

// NewSquare1 creates an identity square-1.
func NewSquare1() Square1 {
	var res Square1

	res.MiddleSquare = true
	res.BottomCount = 8
	res.TopCount = 8

	// Setup identity top layer.
	res.Top[0] = Piece{false, Yellow, Blue, Orange}
	res.Top[1] = Piece{true, Yellow, Orange, 0}
	res.Top[2] = Piece{false, Yellow, Orange, Green}
	res.Top[3] = Piece{true, Yellow, Green, 0}
	res.Top[4] = Piece{false, Yellow, Green, Red}
	res.Top[5] = Piece{true, Yellow, Red, 0}
	res.Top[6] = Piece{false, Yellow, Red, Blue}
	res.Top[7] = Piece{true, Yellow, Blue, 0}

	// Setup identity bottom layer.
	res.Bottom[0] = Piece{true, White, Blue, 0}
	res.Bottom[1] = Piece{false, White, Blue, Red}
	res.Bottom[2] = Piece{true, White, Red, 0}
	res.Bottom[3] = Piece{false, White, Red, Green}
	res.Bottom[4] = Piece{true, White, Green, 0}
	res.Bottom[5] = Piece{false, White, Green, Orange}
	res.Bottom[6] = Piece{true, White, Orange, 0}
	res.Bottom[7] = Piece{true, White, Orange, Blue}

	return res
}

// TurnBottom turns the bottom the next reasonable amount.
// This returns the number of 30 degree clockwise turns.
func (s Square1) TurnBottom() int {
	angle := 0
	for {
		piece := s.Bottom[s.BottomCount-1]
		copy(s.Bottom[1:], s.Bottom[0:])
		s.Bottom[0] = piece
		if piece.Edge {
			angle++
		} else {
			angle += 2
		}
		if s.ValidBottom() {
			return angle
		}
	}
}

// TurnTop turns the top the next reasonable amount.
// This returns the number of 30 degree clockwise turns.
func (s Square1) TurnTop() int {
	angle := 0
	for {
		piece := s.Top[s.TopCount-1]
		copy(s.Top[1:], s.Top[0:])
		s.Top[0] = piece
		if piece.Edge {
			angle++
		} else {
			angle += 2
		}
		if s.ValidTop() {
			return angle
		}
	}
}

// ValidBottom returns true if the bottom can be moved.
func (s Square1) ValidBottom() bool {
	angle := 0
	for i := 0; i < s.BottomCount; i++ {
		x := s.Bottom[i]
		if x.Edge {
			angle++
		} else {
			angle += 2
		}
		if angle == 6 {
			return true
		} else if angle > 6 {
			return false
		}
	}
	return false
}

// ValidTop returns true if the top can be moved.
func (s Square1) ValidTop() bool {
	angle := 0
	for i := 0; i < s.TopCount; i++ {
		x := s.Top[i]
		if x.Edge {
			angle++
		} else {
			angle += 2
		}
		if angle == 6 {
			return true
		} else if angle > 6 {
			return false
		}
	}
	return false
}

// TurnRight turns the right part of the puzzle 180 degrees.
func (s Square1) TurnRight() {
	s.MiddleSquare = !s.MiddleSquare

	// Backup the old state.
	var top [10]Piece
	var bottom [10]Piece
	topCount := s.TopCount
	bottomCount := s.BottomCount
	copy(top[:], s.Top[:])
	copy(bottom[:], s.Bottom[:])

	// Figure out how many pieces are in the unchanged part of the top.
	angle := 0
	for i := 0; i < topCount; i++ {
		piece := top[i]
		if piece.Edge {
			angle++
		} else {
			angle += 2
		}
		if angle >= 6 {
			s.TopCount = i + 1
		}
	}

	// Copy second half of top to bottom.
	s.BottomCount = 0
	for i := s.TopCount; i < topCount; i++ {
		piece := top[i]
		s.Bottom[s.BottomCount] = piece
		s.BottomCount++
	}

	// Copy first half of bottom to second half of top.
	angle = 0
	bottomUsed := 0
	for i := 0; angle < 6; i++ {
		piece := bottom[i]
		s.Top[s.TopCount] = piece
		s.TopCount++
		if piece.Edge {
			angle++
		} else {
			angle += 2
		}
		bottomUsed++
	}

	// Copy the second half of the bottom back to the bottom.
	for i := bottomUsed; i < bottomCount; i++ {
		s.Bottom[s.BottomCount] = bottom[i]
		s.BottomCount++
	}
}
