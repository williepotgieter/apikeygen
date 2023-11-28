package util

import (
	"hash/crc32"
	"regexp"
	"strconv"
	"strings"
)

func ValidateApiKey(key string) (bool, error) {
	var (
		keyParts []string
		secret   []byte
		checksum uint32
	)

	// Check that key contains three parts (label, secret and checksum)
	keyParts = strings.Split(key, "_")
	if len(keyParts) != 3 {
		return false, nil
	}

	// Extract secret
	secret = []byte(keyParts[1])
	// Check that the length of the secret is <= 256 (uint8)
	if len(secret) > 256 {
		return false, nil
	}

	// Extract checksum
	if cks, err := strconv.ParseUint(keyParts[2], 10, 32); err != nil {
		return false, err
	} else {
		checksum = uint32(cks)
	}

	// Check that the secret contains only a-zA-Z0-9
	if match, err := regexp.MatchString("^[0-9a-zA-Z]+$", keyParts[1]); err != nil {
		return false, err
	} else if !match {
		return false, nil
	}

	// Verify the checksum
	if checksum != crc32.ChecksumIEEE(secret) {
		return false, nil
	}

	return true, nil
}
