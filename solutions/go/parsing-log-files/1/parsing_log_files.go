package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	isValidLogLineRegex := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return isValidLogLineRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	doubleLineRgx := regexp.MustCompile(`<([\*=~-]*)>`)
	return doubleLineRgx.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	quotedPasswordRgx := regexp.MustCompile(`(?i)\"[\S\s]*(password)+[\S\s]*\"`)

	sum := 0
	for _, line := range lines {
		if quotedPasswordRgx.MatchString(line) {
			sum += 1
		}
	}
	return sum
}

func RemoveEndOfLineText(text string) string {
	endOfLineRgx := regexp.MustCompile(`end-of-line\d+`)
	return strings.Join(endOfLineRgx.Split(text, -1), "")
}

func TagWithUserName(lines []string) []string {
	findUserRgx := regexp.MustCompile(` User\s+(\w+)`)
	for i, val := range lines {
		res := findUserRgx.FindStringSubmatch(val)
		if res == nil {
			continue
		}
		lines[i] = fmt.Sprintf("[USR] %s %s", res[1], val)

	}
	return lines
}
