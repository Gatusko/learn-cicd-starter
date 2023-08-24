package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	var myHeaders http.Header
	myHeaders = make(http.Header)
	myHeaders.Add("Authorization", "ApiKey aa")
	_, err := GetAPIKey(myHeaders)
	if err != nil {
		t.Fatalf("Should have header authoirzation: %s", err)
	}

}
