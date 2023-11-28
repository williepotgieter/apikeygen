package keymaker

import (
	"fmt"

	"github.com/williepotgieter/keymaker/internal/util"
)

func NewKey(label string, length uint8) (string, error) {
	secret, err := util.GenerateRandomBytes(int(length))
	if err != nil {
		return "", err
	}

	checksum := util.GenerateChecksum(secret)

	key := fmt.Sprintf("%s_%s_%d", label, string(secret), checksum)

	return key, nil
}

func VerifyKey(key string) (bool, error) {
	return util.ValidateChecksum(key)
}
