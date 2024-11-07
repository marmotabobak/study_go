package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func ReturnRandomIntString(max int) string {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(1000))
	return fmt.Sprintf("%s", randInt)
}