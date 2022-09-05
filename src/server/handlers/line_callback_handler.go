package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/go-redis/redis"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/utils/custom_math"
	"github.com/shota-imoto/helixf/lib/utils/domain"
	"github.com/shota-imoto/helixf/lib/utils/line"
	"github.com/shota-imoto/helixf/src/server/supports"
)

func LineCallbackHandler(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(os.Getenv("LINE_SECRET"), os.Getenv("LINE_ACCESS_TOKEN"))

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}

	events, err := bot.ParseRequest(r)

	if err != nil {
		supports.ErrorHandler(w, r, err)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeJoin {
			err := JoinCallback(w, *event)
			if err != nil {
				supports.ErrorHandler(w, r, err)
				break
			}
		}
	}
}

type StateInformation struct {
	State        string `json:"state"`
	RedirectPath string `json:"redirect_path"`
	UrlQuery     string `json:"url_query"`
}

func (info *StateInformation) MarshalBinary() ([]byte, error) {
	json_info, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	return []byte(json_info), nil
}

func (info *StateInformation) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &info)
	if err != nil {
		return err
	}
	return nil
}

func LineAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	auth_url, err := url.Parse(line.LoginURL)

	if err != nil {
		panic(err)
	}

	// generate URL
	query := auth_url.Query()
	query.Set("response_type", "code")
	query.Set("client_id", line.ClientId)
	query.Set("redirect_uri", domain.Url+"/assert_auth")
	query.Set("scope", "profile openid")

	// OAuthのstate作成・保存。加えて、ログイン後にアクセスするリダイレクトURLとクエリも保存する
	state := custom_math.RandLetter(12)
	ctx := context.Background()

	info := StateInformation{State: state, RedirectPath: r.URL.Query().Get("redirect_path"), UrlQuery: r.URL.Query().Get("query")}
	err = db.Kvs.Set(ctx, state, &info, 0).Err()

	fmt.Println(1111111)
	fmt.Println(err)
	fmt.Println(1111111)
	if err != nil {
		panic(err)
	}
	query.Set("state", state)

	auth_url.RawQuery, err = url.QueryUnescape(query.Encode())

	// redirect response
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", auth_url.String())
	w.WriteHeader(http.StatusMovedPermanently)
}

func AssertAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	errs, _ := r.URL.Query()["error"]

	if len(errs) > 0 {
		fmt.Println("auth error: ", errs[0], r.URL.Query()["error_description"])
	}

	code, codeOk := r.URL.Query()["code"]
	state, stateOk := r.URL.Query()["state"]

	if !codeOk || !stateOk {
		fmt.Println("parameter missing: 503レスポンス返す")
	}

	ctx := context.Background()
	var info StateInformation
	err := db.Kvs.Get(ctx, state[0]).Scan(&info)

	if err == redis.Nil { // TODO: valの確認。日付をパースして期限切れチェックを行いたい
		fmt.Println("401返す。railsのCSRF invalidと同じ挙動をしたい")
	}

	if err != nil {
		panic(err)
	}

	client := line.GetAccessTokenClient{AuthorizationCode: code[0]}
	tokenResponse, err := client.Do()

	if err != nil {
		panic(err)
	}

	frontend_url, err := url.Parse(domain.FrontendUrl + "/" + info.RedirectPath + info.UrlQuery)
	query := frontend_url.Query()
	query.Set("authorization", tokenResponse.IdToken)

	frontend_url.RawQuery = query.Encode()
	w.Header().Set("location", frontend_url.String())
	w.WriteHeader(http.StatusMovedPermanently)
}

func JoinCallback(w http.ResponseWriter, event linebot.Event) error {
	message := "下記のURLを開くとWebブラウザから設定が行えるようになります。\n" + domain.FrontendUrl + "/groups?group_id=" + event.Source.GroupID
	bot, err := linebot.New(os.Getenv("LINE_SECRET"), os.Getenv("LINE_ACCESS_TOKEN"))
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	if _, err := bot.PushMessage(event.Source.GroupID, linebot.NewTextMessage(message)).Do(); err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	return nil
}
