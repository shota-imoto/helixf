package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/models/line_model"
	"github.com/shota-imoto/helixf/lib/service/line_service"
	"github.com/shota-imoto/helixf/lib/utils/line"
	"github.com/shota-imoto/helixf/src/server/middleware"
	"github.com/shota-imoto/helixf/src/server/supports"
)

func RegisterGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called Post Groups")

	user := r.Context().Value(middleware.AuthorizationUserKey).(helixf_user.User)

	// ユーザーがグループに所属しているか確認する
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}

	client := line.CheckMenberClient{UserId: user.LineId}
	err = json.Unmarshal(reqBody, &client)

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}
	_, err = client.Do()

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}

	group, err := line_service.FindOrCreateGroupByGroupId(client.GroupId)

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}

	line_service.JoinGroup(group, user)

	w.WriteHeader(200)
	return
}

type GetListGroupsResponse struct {
	Groups []line_model.LineGroup `json:"groups,array"`
}

func GetListGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get list groups")
	user := r.Context().Value(middleware.AuthorizationUserKey).(helixf_user.User)
	fmt.Println(user)

	groups, err := line_service.GetListGroups(user)

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}

	response := GetListGroupsResponse{groups}

	response_json, err := json.Marshal(response)
	fmt.Println(string(response_json))

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}
	w.Write(response_json)
	w.WriteHeader(200)
}
