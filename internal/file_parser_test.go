package internal_test

import (
	"testing"

	"github.com/ilya-skoropad/app-config-provider/internal"
	"github.com/stretchr/testify/assert"
)

const (
	TestFileContent = "APP_NAME=1\nAPP_KEY=\"Ad3!=6^w\"\n#HI=true\n# simple comment\n\n\nLAST_VAR = data"
)

func TestParseFile(t *testing.T) {
	expectedMap := map[string]string{
		"APP_NAME": "1",
		"APP_KEY":  "Ad3!=6^w",
		"HI":       "true",
		"LAST_VAR": "data",
	}

	got, err := internal.ParseFile(TestFileContent)
	if err != nil {
		t.Fail()
		return
	}

	assert.Equal(t, expectedMap, got, "maps should be the same")
}
