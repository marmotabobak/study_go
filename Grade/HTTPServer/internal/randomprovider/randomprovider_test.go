package randomprovider

import (
	"testing"
	"httpserver/internal/utils"
)

func TestRandomProvider(t *testing.T) {
	var sum int

	sign := -1
	rp := NewRandomNumberProvider()

	for i := 0; i <= 10; i++ {
		randomNumber := rp.GetRandomInt()
		randomnumberType := utils.CheckType(randomNumber)
		if randomnumberType != utils.Int {
			t.Errorf("RandomNumber expected to be int, got: %s", randomnumberType)
		}

		sum += sign * randomNumber
		sign *= -1
	}

	if sum == 0 {
		t.Error("RandomNumber expected to provide different random numbers, got same")
	}
}