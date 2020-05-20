package strutils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	SetA = "abcdefghijklmnopqrstuvwxyz"
	SetN = "0123456789"
)

func RandomA(length int) string {
	return Random(length, SetA)
}

func RandomN(length int) string {
	return Random(length, SetN)
}

func RandomANS(length int) string {
	return Random(length, SetA+strings.ToUpper(SetA)+SetN)
}

func RandomANI(length int) string {
	return Random(length, SetA+SetN)
}

func Random(length int, set string) string {
	randomStr := make([]rune, length)

	bigLength := big.NewInt(int64(len(set)))
	for i := 0; i < length; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigLength)
		randomStr[i] = []rune(set)[randomInt.Int64()]
	}

	return string(randomStr)
}
