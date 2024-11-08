package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Type string

const (
	Int Type = "int"
	String Type = "string"
	Unknown Type = "unknown"
)

func ReturnRandomIntString(max int) string {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(1000))
	return fmt.Sprintf("%s", randInt)
}

func CheckType(i interface{}) Type {
	switch i.(type) {
	case int:
		return Int
	case string:
		return String
	default:
		return Unknown
	}
}