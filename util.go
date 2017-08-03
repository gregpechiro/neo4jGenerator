package neo4jGenerator

import (
	"strings"
	"unicode"

	"github.com/gregpechiro/structFields"
)

func ToSnake(s string) string {
	ss := strings.Fields(s)
	s = strings.Join(ss, "-")
	s = strings.ToLower(string(s[0])) + s[1:]
	for i := 0; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			s = s[:i] + "-" + strings.ToLower(string(s[i])) + s[i+1:]
		}
	}
	return s
}

func toLowerFirst(s string) string {
	if len(s) > 1 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	return strings.ToLower(s)
}

func GetSqlType(f structFields.Field) string {
	return ""
}
