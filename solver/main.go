package main

import (
	"fmt"
	"github.com/unixpickle/squareone"
)

func main() {
	sq1 := squareone.Input()
	
	fmt.Println("generating heuristic...")
	h := squareone.MakeShapeHeuristic()
	
	fmt.Println(len(h))
	fmt.Println("predicted depth:", h.Bound(sq1))
	
	fmt.Println("solving shape...")
	for depth := 0; depth < 13; depth++ {
		fmt.Println("depth", depth)
		solution := squareone.RestoreShape(sq1, h, depth)
		if solution != nil {
			fmt.Println(solution)
			break
		}
	}
}
