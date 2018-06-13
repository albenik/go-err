package errx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trimSourcePathPrefix(t *testing.T) {
	f, l := getCallerInfo(1)
	assert.Equal(t, "github.com/albenik/go-errx/caller_test.go", f)
	assert.Equal(t, 10, l)
}
