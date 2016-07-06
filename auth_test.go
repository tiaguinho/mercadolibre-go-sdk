package meli

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	client = Client{
		ClientID:     907054494590799,
		ClientSecret: "x7qFo8AudrLHEsDWm96Kwfu1xbYTiWbW",
	}

	code        = "your code"
	redirectUrl = "redirect url"
)

//Test GetAuthUrl function
func TestGetAuthUrl(t *testing.T) {
	result, _ := client.GetAuthUrl("", AuthUrls["MLB"])

	expected := "https://auth.mercadolivre.com.br/authorization?client_id=907054494590799&redirect_uri=&response_type=code"
	if reflect.DeepEqual(expected, result) == false {
		t.Error(fmt.Sprintf("Expected: %s - Got: %s", expected, result))
	}
}

//Test Authorize function
func TestAuthorize(t *testing.T) {
	err := client.Authorize(code, redirectUrl)

	if err != nil {
		t.Error(err)
	}
}

//Test RefreshAccessToken function
func TestRefreshAccessToken(t *testing.T) {
	err := client.RefreshAccessToken()

	if err != nil {
		t.Error(err)
	}
}
