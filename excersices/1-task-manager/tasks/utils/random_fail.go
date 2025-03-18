package utils

import (
	"errors"
	"math/rand/v2"
)

// simulate random fails;
// chance between 0-100; 100 means on 100% failure
func GetRandomFail(message string, chance int) error {
	max := 100

	if rand.IntN(max) <= chance {
		return errors.New(message)
	}

	return nil
}
