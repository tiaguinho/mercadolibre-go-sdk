package meli

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

const (
	version = "1.1.0"

	apiRootUrl = "https://api.mercadolibre.com"
	oauthUrl   = "oauth/token"
)

var (
	AuthUrls = map[string]string{
		"MLA": "https://auth.mercadolibre.com.ar", // Argentina
		"MLB": "https://auth.mercadolivre.com.br", // Brasil
		"MCO": "https://auth.mercadolibre.com.co", // Colombia
		"MCR": "https://auth.mercadolibre.com.cr", // Costa Rica
		"MEC": "https://auth.mercadolibre.com.ec", // Ecuador
		"MLC": "https://auth.mercadolibre.cl",     // Chile
		"MLM": "https://auth.mercadolibre.com.mx", // Mexico
		"MLU": "https://auth.mercadolibre.com.uy", // Uruguay
		"MLV": "https://auth.mercadolibre.com.ve", // Venezuela
		"MPA": "https://auth.mercadolibre.com.pa", // Panama
		"MPE": "https://auth.mercadolibre.com.pe", // Peru
		"MPT": "https://auth.mercadolibre.com.pt", // Prtugal
		"MRD": "https://auth.mercadolibre.com.do", // Dominicana
	}
)

//getAuthUrl
func (c *Client) GetAuthUrl(redirectUri, authUrl string) (authUri string, err error) {
	if authUrl != "" {
		query := url.Values{}

		query.Set("response_type", "code")
		query.Set("client_id", strconv.Itoa(c.ClientID))
		query.Set("redirect_uri", redirectUri)

		authUri = fmt.Sprintf("%s/authorization?%s", authUrl, query.Encode())
	} else {
		err = errors.New("auth url is empty")
	}

	return
}

//Authorize
func (c *Client) Authorize(code, redirectUri string) (err error) {
	params := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     strconv.Itoa(c.ClientID),
		"client_secret": c.ClientSecret,
		"code":          code,
		"redirect_uri":  redirectUri,
	}

	data, mlErr := execute("POST", oauthUrl, params, nil)

	if mlErr == nil {
		json.Unmarshal(data, &c.MLToken)
	} else {
		err = mlErr
	}

	return
}

//RefreshAccessToken
func (c *Client) RefreshAccessToken() (err error) {
	params := map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     strconv.Itoa(c.ClientID),
		"client_secret": c.ClientSecret,
		"refresh_token": c.MLToken.RefreshToken,
	}

	data, mlErr := execute("POST", oauthUrl, params, nil)

	if mlErr == nil {
		json.Unmarshal(data, &c.MLToken)
	} else {
		err = mlErr
	}

	return
}
