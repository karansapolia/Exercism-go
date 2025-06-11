package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	for index := range cb {
		for _, value := range cb[index] {
			if value == true {
				count++
			}
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	for index := range cb {
		for fileIndex, value := range cb[index] {
			if fileIndex == rank-1 && value == true {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	countSquares := 0
	for index := range cb {
		for range cb[index] {
			countSquares++
		}
	}
	return countSquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	panic("Please implement CountOccupied()")
}
