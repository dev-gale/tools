package tools

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Stringer string

func String(v string) Stringer {
	return Stringer(v)
}

func (s Stringer) String() string {
	return string(s)
}

func (s Stringer) Len() int {
	return len(s)
}

func (s Stringer) Size() int {
	return utf8.RuneCountInString(s.String())
}

func (s Stringer) IsEmpty() bool {
	return s.Len() == 0
}

func (s Stringer) Integer() Integer {
	i, _ := strconv.Atoi(s.String())
	return Integer(i)
}

func (s Stringer) Floater() Floater {
	f, _ := strconv.ParseFloat(s.String(), 64)
	return Floater(f)
}

func (s Stringer) Boolean() Boolean {
	b, _ := strconv.ParseBool(s.String())
	return Boolean(b)
}

func (s Stringer) TrimSpace() Stringer {
	return Stringer(strings.TrimSpace(s.String()))
}

func (s Stringer) Trim(cut string) Stringer {
	return Stringer(strings.Trim(s.String(), cut))
}

func (s Stringer) TrimLeft(cut string) Stringer {
	return Stringer(strings.TrimLeft(s.String(), cut))
}

func (s Stringer) TrimRight(cut string) Stringer {
	return Stringer(strings.TrimRight(s.String(), cut))
}

func (s Stringer) HasSuffix(cut string) bool {
	return strings.HasSuffix(s.String(), cut)
}

func (s Stringer) HasPrefix(cut string) bool {
	return strings.HasPrefix(s.String(), cut)
}

func (s Stringer) Remove(cut string) Stringer {
	before, after, _ := strings.Cut(s.String(), cut)
	return Stringer(before).Add(after)
}

func (s Stringer) Cut(cut string) (string, string, bool) {
	return strings.Cut(s.String(), cut)
}

func (s Stringer) CutSuffix(cut string) Stringer {
	before, _ := strings.CutSuffix(s.String(), cut)
	return Stringer(before)
}

func (s Stringer) CutPrefix(cut string) Stringer {
	after, _ := strings.CutPrefix(s.String(), cut)
	return Stringer(after)
}

func (s Stringer) Add(ss ...string) Stringer {
	if len(ss) == 0 {
		return s
	}
	var str strings.Builder
	str.WriteString(s.String())
	for _, v := range ss {
		str.WriteString(v)
	}
	return Stringer(str.String())
}

func (s Stringer) Added(ss string) Stringer {
	return String(ss).Add(ss)
}

// FirstUpper 首字母大写
func (s Stringer) FirstUpper() Stringer {
	if s.Len() == 0 {
		return s
	}
	runes := []rune(s)
	if unicode.IsLetter(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return Stringer(runes)
}

// FirstLower 首字母小写
func (s Stringer) FirstLower() Stringer {
	if s.Len() == 0 {
		return s
	}
	runes := []rune(s)
	if unicode.IsLetter(runes[0]) {
		runes[0] = unicode.ToLower(runes[0])
	}
	return Stringer(runes)
}

func (s Stringer) LastUpper() Stringer {
	if s.Len() == 0 {
		return s
	}
	runes := []rune(s)
	if unicode.IsLetter(runes[len(runes)-1]) {
		runes[len(runes)-1] = unicode.ToUpper(runes[len(runes)-1])
	}
	return Stringer(runes)
}

func (s Stringer) LastLower() Stringer {
	if s.Len() == 0 {
		return s
	}
	runes := []rune(s)
	if unicode.IsLetter(runes[len(runes)-1]) {
		runes[len(runes)-1] = unicode.ToLower(runes[len(runes)-1])
	}
	return Stringer(runes)
}

func (s Stringer) Upper() Stringer {
	if s.Len() == 0 {
		return s
	}
	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		if unicode.IsLetter(runes[i]) {
			runes[i] = unicode.ToUpper(runes[i])
		}
	}
	return Stringer(runes)
}

func (s Stringer) Lower() Stringer {
	if s.Len() == 0 {
		return s
	}
	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		if unicode.IsLetter(runes[i]) {
			runes[i] = unicode.ToLower(runes[i])
		}
	}
	return Stringer(runes)
}

// Camel2Snake 驼峰转蛇形
func (s Stringer) Camel2Snake() Stringer {
	runes := []rune(s.TrimSpace())
	var ns strings.Builder
	for _, r := range runes {
		if unicode.IsUpper(r) {
			ns.WriteString("_")
		}
		ns.WriteRune(unicode.ToLower(r))
	}
	return Stringer(ns.String()).Trim("_")
}

func (s Stringer) Snake2BigCamel() string {
	runes := []rune(s.TrimSpace())

	if len(runes) == 0 {
		return ""
	}

	if len(runes) == 1 {
		return string(unicode.ToUpper(runes[0]))
	}

	var ns strings.Builder
	if runes[0] != '_' {
		ns.WriteRune(runes[0])
	}
	for i := 1; i < len(runes); i++ {
		if runes[i] == '_' {
			continue
		}

		if runes[i-1] == '_' {
			ns.WriteRune(unicode.ToUpper(runes[i]))
		} else {
			ns.WriteRune(unicode.ToLower(runes[i]))
		}
	}
	return Stringer(ns.String()).FirstUpper().String()
}

func (s Stringer) Snake2LittleCamel() string {
	runes := []rune(s.TrimSpace())

	if len(runes) == 0 {
		return ""
	}

	if len(runes) == 1 {
		return string(unicode.ToLower(runes[0]))
	}

	var ns strings.Builder
	if runes[0] != '_' {
		ns.WriteRune(runes[0])
	}
	for i := 1; i < len(runes); i++ {
		if runes[i] == '_' {
			continue
		}

		if runes[i-1] == '_' {
			ns.WriteRune(unicode.ToUpper(runes[i]))
		} else {
			ns.WriteRune(unicode.ToLower(runes[i]))
		}
	}
	return Stringer(ns.String()).FirstLower().String()
}

func (s Stringer) Split(sep string) []string {
	return strings.Split(string(s), sep)
}
