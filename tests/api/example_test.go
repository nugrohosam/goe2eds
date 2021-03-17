package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	factories "github.com/nugrohosam/goe2eds/tests/factories"
	utilities "github.com/nugrohosam/goe2eds/tests/utilities"
	viper "github.com/spf13/viper"
	assert "github.com/stretchr/testify/assert"
)

// TestRun ...
func TestRun(t *testing.T) {
	InitialTest(t)
	defer utilities.DbCleaner(t)
}