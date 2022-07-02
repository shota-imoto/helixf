package line

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/shota-imoto/helixf/lib/utils/domain"
)

var LoginURL string = "https://access.line.me/oauth2/v2.1/authorize"
var TokenURL string = "https://api.line.me/oauth2/v2.1/token"
var ClientId string = "1657099914"

func ParseIdToken(idToken string) (*jwt.Token, error) {
	tokenString := idToken
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("LINE_LOGIN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// Line認証により取得したAuthorizationCodeを元に、AccessTokenをリクエストするClient
type GetAccessTokenClient struct {
	AuthorizationCode string
}

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpireIn     int    `json:"expire_in"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type LineLoginErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	State            string `json:"state"`
}

func (c *GetAccessTokenClient) Do() (*AccessTokenResponse, error) {
	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Set("code", c.AuthorizationCode)
	values.Set("redirect_uri", domain.Url+"/assert_auth")
	values.Set("client_id", ClientId)
	values.Set("client_secret", os.Getenv("LINE_LOGIN_SECRET"))

	encodedValues, err := url.QueryUnescape(values.Encode())

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", TokenURL, strings.NewReader(encodedValues))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := new(http.Client)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	var tokenResponse AccessTokenResponse
	resBody, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(resBody, &tokenResponse)

	if err != nil {
		return nil, err
	}

	if tokenResponse.AccessToken == "" {
		var errorResponse LineLoginErrorResponse
		err = json.Unmarshal(resBody, &errorResponse)
		return nil, errors.New(fmt.Sprintf("error: %v, message: %v", errorResponse.Error, errorResponse.ErrorDescription))
	}

	return &tokenResponse, nil
}
