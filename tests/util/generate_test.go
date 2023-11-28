package util_test

import (
	"testing"

	"github.com/williepotgieter/keymaker/internal/util"
)

func Test_GenerateRandomBytes(t *testing.T) {
	n := 32

	// ASCII characters a-z => 97-122 / 65-90
	// ASCII digits 0-9 => 48-57
	isAllowed := func(b byte) bool {
		return (b >= 97 && b <= 122) || (b >= 65 && b <= 90) || (b >= 48 && b <= 57)
	}

	result, err := util.GenerateRandomBytes(n)
	if err != nil {
		t.Fatalf("Error when calling GenerateRandomString(%d)", n)
	}

	// Must be of correct length
	if len(result) != n {
		t.Errorf("Expected result to contain %d bytes, but it contains %d bytes", n, len(result))
	}

	// Only allow a-z, A-Z and 0-9
	for i := 0; i < n; i++ {
		if !isAllowed(result[i]) {
			t.Errorf("Expected all bytes to represent a-zA-Z0-9, but got %s, which contains \"%s\"", string(result), string(result[i]))
		}
	}
}
