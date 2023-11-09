package keymaker

import (
	"errors"
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Keymaker interface {
	GenerateApiKey(label string) (apiKey string, err error)
	VerifyApiKey(apiKey string) (err error)
}

type adapter struct{}

func New() Keymaker {
	return &adapter{}
}

// GenerateApiKey returns an api key in the following format: <LABEL>_<KEY>_<CHECKSUM>
// Examples
// acme_a969bb27749049b3a222215e0e7b33dc_3901431202
func (a *adapter) GenerateApiKey(label string) (apiKey string, err error) {

	randomUuid, err := uuid.NewRandom()
	if err != nil {
		return
	}

	key := strings.ReplaceAll(randomUuid.String(), "-", "")

	hasher := crc32.NewIEEE()
	hasher.Write([]byte(key))

	checksum := hasher.Sum32()

	apiKey = fmt.Sprintf("%s_%s_%d", label, key, checksum)

	return
}

func (a *adapter) VerifyApiKey(apiKey string) (err error) {
	// Check that api key is not empty
	if apiKey == "" {
		err = errors.New("empty value provided for api key")
		return
	}

	keyParts := strings.Split(apiKey, "_")

	// Check that api key has three parts
	if len(keyParts) != 3 {
		err = errors.New("invalid api key format")
		return
	}

	// Organize the key parts
	key, checksumStr := keyParts[1], keyParts[2]

	checksum, err := strconv.ParseUint(checksumStr, 10, 32)
	if err != nil {
		err = errors.Join(err, errors.New("error while parsing api key"))
		return
	}

	// Validate checksum
	crc32Checksum := crc32.ChecksumIEEE([]byte(key))

	if crc32Checksum != uint32(checksum) {
		err = errors.New("invalid checksum for api key")
	}

	return
}
