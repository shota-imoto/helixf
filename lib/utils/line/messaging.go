package line

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/utils/helixf_env"
)

// TODO: ディレクトリ・ファイル名全般に違和感のため再検討。APIやAPI clientのような名称が含まれるべき

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

func LinebotClient() (*linebot.Client, error) {
	return linebot.New(os.Getenv("LINE_SECRET"), os.Getenv("LINE_ACCESS_TOKEN"))
}

func (client CheckMenberClient) CheckMemberURL() string {
	return "https://api.line.me/v2/bot/group/" + client.GroupId + "/member/" + client.UserId
}

func (client CheckMenberClient) Do() (CheckMemberResponse, error) {
	fmt.Println("Called check member url request: ", client)
	req, err := http.NewRequest("GET", client.CheckMemberURL(), nil)

	if err != nil {
		return CheckMemberResponse{}, nil
	}
	req.Header.Set("Authorization", "Bearer "+helixf_env.ChannelAccessToken)
	c := new(http.Client)
	res, err := c.Do(req)

	if err != nil {
		return CheckMemberResponse{}, nil
	}
	var response CheckMemberResponse
	reqBody, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(reqBody, &response)

	if err != nil {
		return CheckMemberResponse{}, nil
	}

	if response.Message != "" {
		return CheckMemberResponse{}, errors.New(response.Message)
	}

	return response, nil
}

type Messager struct {
	GroupId         string
	RegularSchedule regular_schedule.RegularSchedule
	Wrapper         LineWrapper
}

// テストでDIする用のinterface。
type LineWrapper interface {
	PushMessage(string, ...linebot.SendingMessage) (string, error)
}

type LineBotWrapper struct {
	Bot *linebot.Client
}

func (wrapper *LineBotWrapper) PushMessage(id string, messages ...linebot.SendingMessage) (string, error) {
	res, err := wrapper.Bot.PushMessage(id, messages...).Do()

	if err != nil {
		return "", err
	}

	return res.RequestID, nil
}

// テスト用。
type DummyBotWrapper struct {
	PushedId       string
	PushedMessages []string
	CalledCount    int
}

func (wrapper *DummyBotWrapper) PushMessage(id string, messages ...linebot.SendingMessage) (string, error) {
	wrapper.PushedId = id

	for _, message := range messages {
		switch message.(type) {
		case *linebot.TemplateMessage:
			wrapper.PushedMessages = append(wrapper.PushedMessages, TextFromTemplate(message.(*linebot.TemplateMessage).Template))
			wrapper.CalledCount++
		default:
			fmt.Println("undefined in helixf")
		}
	}

	return "", nil
}

func TextFromTemplate(t linebot.Template) string {
	switch t.(type) {
	case *linebot.ConfirmTemplate:
		return t.(*linebot.ConfirmTemplate).Text
	default:
		return "undefined template in helixf"
	}
}

func (messager *Messager) SendConfirm() error {
	// https://developers.line.biz/en/reference/messaging-api/#confirm
	leftBtn := linebot.NewMessageAction("accept", "accept")
	rightBtn := linebot.NewMessageAction("deny", "deny")

	// 定期スケジュールの日付と出欠確認の文書
	template := linebot.NewConfirmTemplate("Can you attend schedule at "+messager.RegularSchedule.Label(), leftBtn, rightBtn)
	message := linebot.TemplateMessage{AltText: "attend confirmation", Template: template}

	_, err := messager.Wrapper.PushMessage(messager.GroupId, &message)

	if err != nil {
		return err
	}

	return nil
}
