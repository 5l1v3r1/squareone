package squareone

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Input() Square1 {
	res := NewSquare1()

	fmt.Print("Is the middle square? [y/n]: ")
	res.MiddleSquare = (inputLine() == "y")

	/*
	fmt.Println("")
	fmt.Println("Hold orange left and blue front so the logo is on the left.")
	fmt.Println("")
	fmt.Println("Turn the top and bottom as needed so the front crack is not" +
		" blocked by a corner piece.")
	fmt.Println("")
	fmt.Println("Enter the top face clockwise starting from the left of the" +
		" crack. For every piece, type either two or three letters depending" +
		" on the type of piece. Type an empty string to finish.")
	fmt.Println("")
	*/
	fmt.Println("Enter front face. Enter blank piece when done.")
	res.TopCount = 0
	for i := 0; i < 10; i++ {
		fmt.Print("Piece: ")
		pieceInfo := inputLine()
		if pieceInfo == "" {
			break
		}
		piece, err := parsePiece(pieceInfo)
		if err != nil {
			fmt.Println("Error:", err)
			i--
			continue
		}
		res.Top[res.TopCount] = *piece
		res.TopCount++
	}
	fmt.Println("")
	fmt.Println("Enter bottom face the same way.")
	fmt.Println("")
	res.BottomCount = 0
	for i := 0; i < 10; i++ {
		fmt.Print("Piece: ")
		pieceInfo := inputLine()
		if pieceInfo == "" {
			break
		}
		piece, err := parsePiece(pieceInfo)
		if err != nil {
			fmt.Println("Error:", err)
			i--
			continue
		}
		res.Bottom[res.BottomCount] = *piece
		res.BottomCount++
	}

	return res
}

func inputLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func parseColor(name uint8) (Color, error) {
	switch name {
	case 'y':
		return Yellow, nil
	case 'w':
		return White, nil
	case 'b':
		return Blue, nil
	case 'g':
		return Green, nil
	case 'r':
		return Red, nil
	case 'o':
		return Orange, nil
	default:
		return 0, errors.New("invalid color: " + string(name))
	}
}

func parsePiece(piece string) (*Piece, error) {
	if len(piece) < 2 || len(piece) > 3 {
		return nil, errors.New("invalid piece: " + piece)
	}
	var res Piece
	var err error
	res.Edge = (len(piece) == 2)
	res.TopColor, err = parseColor(piece[0])
	if err != nil {
		return nil, err
	}
	res.SecondColor, err = parseColor(piece[1])
	if err != nil {
		return nil, err
	}
	if len(piece) == 2 {
		return &res, nil
	}
	res.ThirdColor, err = parseColor(piece[2])
	if err != nil {
		return nil, err
	}
	return &res, nil
}
