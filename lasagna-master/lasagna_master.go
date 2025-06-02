package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodlesQuantity := 0
	sauceQuantity := 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodlesQuantity += 50
		} else if layers[i] == "sauce" {
			sauceQuantity += 0.2
		}
	}

	return noodlesQuantity, sauceQuantity
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	friendsLastIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = friendsLastIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	amountsForXPortions := make([]float64, len(amounts))
	for i := 0; i < len(amounts); i++ {
		amountsForXPortions[i] = (amounts[i] / 2.0) * float64(portions)
	}
	return amountsForXPortions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
