package util

import (
	"crypto/rand"
	"hash/crc32"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	var (
		result           []byte
		randomCharacters []byte
		randomDigits     []byte
		match            bool
		err              error
	)

	isCharacter := func(code byte) bool {
		return (code >= 97 && code <= 122) || (code >= 65 && code <= 90)
	}

	isDigit := func(code byte) bool {
		return code >= 48 && code <= 57
	}

	makeRandomBytes := func(fn func(code byte) bool) ([]byte, error) {
		var (
			randomBytes = []byte{}
			tempBytes   []byte
		)

		for i := 0; i < n; i++ {
			tempBytes = make([]byte, 1)
			match = false

			for !match {
				if _, err = rand.Read(tempBytes); err != nil {
					return nil, err
				} else if fn(tempBytes[0]) {
					match = true
				}
			}

			randomBytes = append(randomBytes, tempBytes[0])
		}

		return randomBytes, nil
	}

	// ASCII characters a-z => 97-122 / 65-90
	// ASCII digits 0-9 => 48-57

	// Generate slize of random letters
	if randomCharacters, err = makeRandomBytes(isCharacter); err != nil {
		return nil, err
	}

	// Generate slice of random digits
	if randomDigits, err = makeRandomBytes(isDigit); err != nil {
		return nil, err
	}

	// Combine the random letters and random digit slices into a
	// new list so that a letter is always followed by a digit
	for i := 0; i < n; i++ {
		result = append(result, randomCharacters[i], randomDigits[i])
	}

	// Shuffle the combined list
	if result, err = shuffle(result); err != nil {
		return nil, err
	}

	// Choose only the first n characters in the list
	result = result[:n]

	return result, nil
}

func GenerateChecksum(b []byte) uint32 {
	hasher := crc32.NewIEEE()
	hasher.Write(b)

	return hasher.Sum32()
}
