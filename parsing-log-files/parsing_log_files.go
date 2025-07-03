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
	count := 0
	re, _ := regexp.Compile(`"[^"]*\b(?i:password)\b[^"]*"`)
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		count += len(matches)
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
