import (
    "strings"
)

func lengthOfLastWord(s string) int {
    stringSplit := strings.Split(s, " ")

	for i := len(stringSplit) - 1; i >= 0; i-- {
		if stringSplit[i] != " " && stringSplit[i] != "" {
			return len(stringSplit[i])
		}
	}
	return 0
}
