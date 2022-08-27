package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/service/line_service"
	"github.com/shota-imoto/helixf/lib/service/regular_schedule_service"
	"github.com/shota-imoto/helixf/lib/utils/domain"
	"github.com/shota-imoto/helixf/src/server/middleware"
	"github.com/shota-imoto/helixf/src/server/supports"
)

type GetListRegularScheduleTemplatesResponse struct {
	RegularScheduleTemplates []regular_schedule.RegularScheduleTemplate `json:"regular_schedule_templates,array"`
}

func GetListRegularScheduleTemplates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH")
	w.Header().Set("Access-Control-Allow-Origin", domain.FrontendUrl)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
	user, _ := r.Context().Value(middleware.AuthorizationUserKey).(helixf_user.User)
	vars := mux.Vars(r)

	group, err := line_service.GetGroupWithTemplate(user, vars["id"])

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}

	response := GetListRegularScheduleTemplatesResponse{group.RegularScheduleTemplates}
	response_json, err := json.Marshal(response)

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}
	w.Write(response_json)
	w.WriteHeader(200)
}

func PostRegularScheduleTemplateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH")
	w.Header().Set("Access-Control-Allow-Origin", domain.FrontendUrl)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
	user, _ := r.Context().Value(middleware.AuthorizationUserKey).(*helixf_user.User)
	fmt.Println(user)

	// curl -X POST http://localhost:8080/regular_schedule_template -H "Content-Type: application/json" -d '{"hour":13, "weekday":5, "week":2}'
	var template regular_schedule.RegularScheduleTemplate
	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &template)
	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}
	template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}

	response, err := json.Marshal(template)

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}
	w.Write(response)
}

func DeleteRegularScheduleTemplateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Origin", domain.FrontendUrl)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")

	user, _ := r.Context().Value(middleware.AuthorizationUserKey).(helixf_user.User)

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		supports.ErrorHandler(w, r, err)
	}

	err = regular_schedule_service.DeleteById(uint(id), user)

	if err != nil {
		supports.ErrorHandler(w, r, err)
	}

	w.WriteHeader(204)
}
