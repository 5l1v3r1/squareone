package squareone

import "strconv"

type Move struct {
	Right  bool
	Top    int
	Bottom int
}

func (m Move) String() string {
	top := m.Top
	if top > 6 {
		top -= 12
	}
	bottom := m.Bottom
	if bottom > 6 {
		bottom -= 6
	}
	ts := strconv.Itoa(top)
	bs := strconv.Itoa(bottom)
	if m.Right {
		return "(" + ts + ", " + bs + ") /"
	} else {
		return "(" + ts + ", " + bs + ")"
	}
}

func RestoreShape(start Square1, h ShapeHeuristic, depth int) []Move {
	if IsCube(start) {
		return []Move{}
	} else if depth == 0 {
		return nil
	} else if h.Bound(start) > depth {
		return nil
	}
	
	topAngle := 0
	for topAngle < 12 {
		if topAngle == 0 && !start.ValidTop() {
			topAngle = start.TurnTop()
		}
		
		bottomAngle := 0
		copy := start
		for bottomAngle < 12 {
			if bottomAngle == 0 && !copy.ValidBottom() {
				bottomAngle = copy.TurnBottom()
			}
			
			if IsCube(copy) {
				return []Move{Move{false, topAngle, bottomAngle}}
			} else {
				// Do a right turn and recurse.
				nester := copy
				nester.TurnRight()
				moves := RestoreShape(nester, h, depth-1)
				if moves != nil {
					return append([]Move{Move{true, topAngle, bottomAngle}},
						moves...)
				}
			}
			
			bottomAngle += copy.TurnBottom()
		}
		
		topAngle += start.TurnTop()
	}
	return nil
}
