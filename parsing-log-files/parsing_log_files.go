package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	validLine, _ := regexp.Compile(`^\[TRC\[|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
	return validLine.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
