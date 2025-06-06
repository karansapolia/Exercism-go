package raindrops

import "strconv"

func Convert(number int) string {
	returnStr := ""
	if number%3 == 0 {
		returnStr += "Pling"
	}
	if number%5 == 0 {
		returnStr += "Plang"
	}
	if number%7 == 0 {
		returnStr += "Plong"
	}
	if returnStr == "" {
		returnStr = strconv.Itoa(number)
	}
	return returnStr
}
