package internal_test

import (
	"testing"

	"github.com/ilya-skoropad/app-config-provider/internal"
	"github.com/stretchr/testify/assert"
)

const (
	TestFileContent = "APP_NAME=1" + internal.ENDL + "APP_KEY=Ad3!=6^w" + internal.ENDL + "#HI=true" + internal.ENDL + "# simple comment" + internal.ENDL + internal.ENDL + "LAST_VAR=SMTH" + internal.ENDL
)

func TestParseFile(t *testing.T) {
	expectedMap := map[string]string{
		"APP_NAME": "1",
		"APP_KEY":  "Ad3!=6^w",
		"LAST_VAR": "SMTH",
	}

	got, err := internal.ParseFile(TestFileContent)
	if err != nil {
		t.Fail()
		return
	}

	assert.Equal(t, expectedMap, got, "maps should be the same")
}
