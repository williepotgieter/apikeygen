package keymaker

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/williepotgieter/keymaker/internal/util"
)

type ApiKey struct {
	Label    string
	Secret   string
	Checksum uint32
}

var (
	ErrInvalidApiKey     = errors.New("invalid api key")
	ErrGenerateNewApiKey = errors.New("error while generating new api key")
	ErrParseApiKey       = errors.New("error while parsing api key")
	ErrValidateApiKey    = errors.New("error while validating api key")
)

// NewApiKey generates a new API key in the format <LABEL>_<SECRET>_<CHECKSUM>
// Example: skp_15nd14iju19qV3s63379a5iJz1jGunI_995215320
// The length (number of characters) of the secret can be set by the 'length' argument.
// Although highly unlikely, it will return the error wrapped with "ErrGenerateNewApiKey" if
// there was a problem during the creation of the key.
func NewApiKey(label string, length uint8) (ApiKey, error) {
	secret, err := util.GenerateRandomBytes(int(length))
	if err != nil {
		return ApiKey{}, errors.Join(ErrGenerateNewApiKey, err)
	}

	checksum := util.GenerateChecksum(secret)

	apiKey := ApiKey{
		Label:    label,
		Secret:   string(secret),
		Checksum: checksum,
	}

	return apiKey, nil
}

// ParseApiKey converts an API key string to an API struct if the string is a valid API key.
// Parsing errors will be wrapped by either "ErrValidateApiKey" or "ErrInvalidApiKey".
func ParseApiKey(key string) (ApiKey, error) {
	var (
		label    string
		secret   string
		checksum uint32
	)

	if valid, err := util.ValidateApiKey(key); err != nil {
		return ApiKey{}, errors.Join(ErrValidateApiKey, err)
	} else if !valid {
		return ApiKey{}, ErrInvalidApiKey
	}

	keyParts := strings.Split(key, "_")

	label = keyParts[0]
	secret = keyParts[1]

	if cks, err := strconv.ParseUint(keyParts[2], 10, 32); err != nil {
		return ApiKey{}, errors.Join(ErrValidateApiKey, err)
	} else {
		checksum = uint32(cks)
	}

	return ApiKey{
		Label:    label,
		Secret:   secret,
		Checksum: checksum,
	}, nil
}

// ValidateApiKey validates a given API key by checking that it is of the correct form <LABEL>_<SECRET>_<CHECKSUM>
// and that the checksum is valid. This function can be integrated into http routing middleware to quickly validate API keys
// from incomming requests and reject invalid API keys as early as possible.
func (k ApiKey) ValidateApiKey(key string) (bool, error) {
	return util.ValidateApiKey(key)
}

// String returns the string value of the API key in the format <LABEL>_<SECRET>_<CHECKSUM>
// eg. skp_15nd14iju19qV3s63379a5iJz1jGunI_995215320
func (k ApiKey) String() string {
	return fmt.Sprintf("%s_%s_%d", k.Label, k.Secret, k.Checksum)
}
