package line

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shota-imoto/helixf/lib/utils/helixf_env"
)

func GroupURL(group_id string) string {
	return "https://api.line.me/v2/bot/group/" + group_id + "/summary"
}

type CheckMemberResponse struct {
	DisplayName string
	UserId      int
	PictureUrl  string
	Message     string
}

type ErrorResponse struct {
	Message string
}

type CheckMenberClient struct {
	AccessToken string `json:"access_token"`
	GroupId     string `json:"group_id"`
	UserId      string
}

func (client CheckMenberClient) CheckMemberURL() string {
	return "https://api.line.me/v2/bot/group/" + client.GroupId + "/member/" + client.UserId
}

func (client CheckMenberClient) Do() (CheckMemberResponse, error) {
	req, err := http.NewRequest("GET", client.CheckMemberURL(), nil)

	if err != nil {
		return CheckMemberResponse{}, nil
	}
	req.Header.Set("Authorization", "Bearer "+helixf_env.ChannelAccessToken)
	c := new(http.Client)
	res, err := c.Do(req)
	fmt.Println(0)

	if err != nil {
		return CheckMemberResponse{}, nil
	}
	var response CheckMemberResponse
	reqBody, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(reqBody, &response)

	if err != nil {
		return CheckMemberResponse{}, nil
	}
	fmt.Println(22)
	fmt.Println("Bearer " + helixf_env.ChannelAccessToken)

	if response.Message != "" {
		return CheckMemberResponse{}, errors.New(response.Message)
	}
	fmt.Println(33)

	return response, nil
}
