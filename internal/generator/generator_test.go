package generator

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOutputPackageName_Success(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	result, err := generateOutputPackageName(dir)
	assert.NoError(t, err)
	assert.Equal(t, "github.com/slasyz/openrest/internal/generator", result)
}

func TestGenerateOutputPackageName_Error(t *testing.T) {
	dir := "/tmp/blabla"

	_, err := generateOutputPackageName(dir)
	assert.ErrorIs(t, err, errNoGoMod)
}
