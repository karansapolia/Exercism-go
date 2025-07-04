package thefarm

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numOfCows int) (float64, error) {
	fodderAmount, FAError := fc.FodderAmount(numOfCows)
	fatteningFactor, FFError := fc.FatteningFactor()
	totalFodder := fodderAmount * fatteningFactor
	if FAError != nil {
		return 0, FAError
	} else if FFError != nil {
		return 0, FFError
	}

	return totalFodder / float64(numOfCows), nil

}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numOfCows int) (float64, error) {
	panic("AAAAAAAA!")
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(noOfCows int) error {
	panic("HAHAHAHHA!")
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
