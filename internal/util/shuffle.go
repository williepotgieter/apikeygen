package util

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func shuffle(list []byte) ([]byte, error) {
	var (
		randIndex int
		err       error
	)

	if len(list) < 2 {
		return nil, errors.New("lengths of list to be shuffled must be >= 2")
	}

	generateRandomNumber := func(max int) (int, error) {
		var (
			rnd    *big.Int
			genErr error
		)
		if max < 1 {
			return 0, errors.New("the maximum number set for a random number must be > 0")
		}

		if rnd, genErr = rand.Int(rand.Reader, big.NewInt(int64(max))); genErr != nil {
			return 0, genErr
		}

		return int(rnd.Int64()), nil
	}

	for lastIndex := len(list) - 1; lastIndex > 0; lastIndex-- {
		if randIndex, err = generateRandomNumber(lastIndex); err != nil {
			return nil, err
		}

		list[lastIndex], list[randIndex] = list[randIndex], list[lastIndex]
	}

	return list, nil
}
