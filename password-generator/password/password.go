package password

import (
	"time"

	"math/rand"
)

const (
	HasSmallAlph = "abcdefghijklmnopqrstuvwxyz"
	HasCapAlph   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	HasNum       = "0123456789"
	HasSymbol    = "!@#$%^&*"
)

type Conditions struct {
	HasSmallAlph  bool
	HasCapAlph    bool
	HasNum        bool
	HasSymbol     bool
	MinLength     int8
	MaxLength     int8
	NumConditions int8
}

var source map[bool]map[bool]map[bool]map[bool]string
var bools = map[bool]struct{}{
	false: {},
	true:  {},
}

func init() {
	source = make(map[bool]map[bool]map[bool]map[bool]string, 2)
	for a := range bools {
		source[a] = make(map[bool]map[bool]map[bool]string, 2)
		for A := range bools {
			source[a][A] = make(map[bool]map[bool]string, 2)
			for num := range bools {
				source[a][A][num] = make(map[bool]string, 2)
				for symb := range bools {
					temp := ""
					if a {
						temp += HasSmallAlph
					}
					if A {
						temp += HasCapAlph
					}
					if num {
						temp += HasNum
					}
					if symb {
						temp += HasSymbol
					}
					source[a][A][num][symb] = temp
				}
			}
		}
	}
}

func CreatePassword(conditions Conditions) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var length int
	if conditions.MinLength == conditions.MaxLength {
		length = int(conditions.MinLength)
	} else {
		length = rand.Intn(int(conditions.MaxLength-conditions.MinLength)) + int(conditions.MinLength)
	}

	passwordByte := make([]byte, 0, length)

	charSet := source[conditions.HasSmallAlph][conditions.HasCapAlph][conditions.HasNum][conditions.HasSymbol]
	sourceLen := len(charSet)

	var (
		smallAlphDone bool = !conditions.HasSmallAlph
		capAlphDone   bool = !conditions.HasCapAlph
		numDone       bool = !conditions.HasNum
		symbolDone    bool = !conditions.HasSymbol
		allDone       bool
	)

	for range length {
		passwordByte = append(passwordByte, charSet[rand.Intn(sourceLen)])
	}
	for _, c := range passwordByte {
		switch {
		case c >= 'a' && c <= 'z':
			if !smallAlphDone {
				smallAlphDone = true
			}
		case c >= 'A' && c <= 'Z':
			if !capAlphDone {
				capAlphDone = true
			}
		case c >= '0' && c <= '9':
			if !numDone {
				numDone = true
			}
		default:
			if !symbolDone {
				symbolDone = true
			}
		}
		if smallAlphDone && capAlphDone && numDone && symbolDone {
			allDone = true
			break
		}
	}
	if !allDone {
		i := 0
		if !smallAlphDone {
			passwordByte[i] = HasSmallAlph[rand.Intn(len(charSet))]
			i++
		}
		if !capAlphDone {
			passwordByte[i] = HasCapAlph[rand.Intn(len(charSet))]
			i++
		}
		if !numDone {
			passwordByte[i] = HasNum[rand.Intn(len(charSet))]
			i++
		}
		if !symbolDone {
			passwordByte[i] = HasSymbol[rand.Intn(len(charSet))]
			i++
		}
	}

	return string(passwordByte)
}
