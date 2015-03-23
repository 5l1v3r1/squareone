package squareone

type ShapeHeuristic map[string]int

func (h ShapeHeuristic) Bound(s Square1) int {
	if x, ok := h[Shape(s)]; ok {
		return x
	} else {
		return 0
	}
}

func MakeShapeHeuristic() ShapeHeuristic {
	nodes := []searchNode{searchNode{NewSquare1(), 0}}
	heuristic := ShapeHeuristic{}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		shape := Shape(node.s)
		if _, ok := heuristic[shape]; ok {
			continue
		}
		heuristic[shape] = node.d
		
		// Branch off.
		for angle := 0; angle < 12; angle += node.s.TurnTop() {
			copy := node.s
			for ba := 0; ba < 12; ba += copy.TurnBottom() {
				if angle == 0 && ba == 0 {
					continue
				}
				nodes = append(nodes, searchNode{copy, node.d+1})
				copy.TurnRight()
				nodes = append(nodes, searchNode{copy, node.d+1})
			}
		}
	}
	return heuristic
}

func Shape(s Square1) string {
	res := ""
	for i := 0; i < s.TopCount; i++ {
		if s.Top[i].Edge {
			res += "E"
		} else {
			res += "C"
		}
	}
	res += " "
	for i := 0; i < s.BottomCount; i++ {
		if s.Bottom[i].Edge {
			res += "E"
		} else {
			res += "C"
		}
	}
	if s.MiddleSquare {
		res += " S"
	} else {
		res += " N"
	}
	return res
}

func IsCube(s Square1) bool {
	return Shape(s) == "CECECECE ECECECEC S"
}

type searchNode struct {
	s Square1
	d int
}
