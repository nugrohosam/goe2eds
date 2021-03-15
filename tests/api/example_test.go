package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	factories "github.com/nugrohosam/gocashier/tests/factories"
	utilities "github.com/nugrohosam/gocashier/tests/utilities"
	viper "github.com/spf13/viper"
	assert "github.com/stretchr/testify/assert"
)

// AuthTestRun ...
func AuthTestRun(t *testing.T) {
	InitialTest(t)
	defer utilities.DbCleaner(t)

	user = factories.CreateUser()

	t.Log("=======>>>> START <<<<======")
	test1(t)
	t.Log("=======>>>>  END  <<<<======")
}

func test1(t *testing.T) {
	url := viper.GetString("app.url")
	port := viper.GetString("app.port")

	data, err := json.Marshal(map[string]interface{}{
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
	},
	)

	if err != nil {
		t.Error(err.Error())
	}

	reader := bytes.NewBuffer(data)
	endpoint := "http://" + url + ":" + port + "/v1/auth/register"

	t.Log("Test Positive Register")
	resp := PerformRequest(Routes, "POST", endpoint, "application/json", reader)
	assert.Equal(t, http.StatusOK, resp.Code)
}
