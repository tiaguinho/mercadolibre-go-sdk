package meli

//Client struct
type Client struct {
	GrantType    string `json:"grant_type,omitempty"`
	Code         string `json:"code"`
	ClientID     int    `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
	MLToken      MLToken
}

//MLToken struct
type MLToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	Scope        string `json:"scope,omitempty"`
	UserId       int    `json:"user_id,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

//MLError struct
type MLError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}

//New
func New(clientID int, clientSecret string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

//NewWithAccessToken
func NewWithAccessToken(clientID int, clientSecret, accessToken, refreshToken string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		MLToken: MLToken{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}
