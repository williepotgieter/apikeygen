package keymaker_test

import (
	"log"
	"regexp"
	"testing"
	"unicode/utf8"

	"github.com/williepotgieter/keymaker/internal/util"
)

func Test_GenerateRandomString(t *testing.T) {
	var (
		stringLength = 32
		resultLength int
	)
	pattern := `[a-z0-9]+`

	result, err := util.GenerateRandomString(stringLength)
	if err != nil {
		log.Println(err)
		t.Fatalf("Error when calling GenerateRandomString(%d)", stringLength)
	}

	resultLength = utf8.RuneCountInString(result)

	if match, err := regexp.MatchString(pattern, result); err != nil {
		t.Fatalf("Error while checking whether %s matches regex pattern %s", result, pattern)
	} else if !match {
		t.Errorf("Expected %s to match regex pattern \"%s\", but it does not", result, pattern)
	} else if resultLength != stringLength {
		t.Errorf("Expected the length of %s to be %d characters, but it is %d characters", result, stringLength, resultLength)
	}
}
