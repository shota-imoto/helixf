package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule_template"
	"github.com/shota-imoto/helixf/lib/service/regular_schedule_service"
	"github.com/shota-imoto/helixf/lib/utils/domain"
	"github.com/shota-imoto/helixf/src/server/middleware"
	"github.com/shota-imoto/helixf/src/server/supports"
)

func PostRegularScheduleTemplateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH")
	w.Header().Set("Access-Control-Allow-Origin", domain.FrontendUrl)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
	user, _ := r.Context().Value(middleware.AuthorizationUserKey).(*helixf_user.User)
	fmt.Println(user)

	// curl -X POST http://localhost:8080/regular_schedule_template -H "Content-Type: application/json" -d '{"hour":13, "weekday":5, "week":2}'
	var template regular_schedule_template.RegularScheduleTemplate
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &template)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Println("unmarshal error")
		supports.ErrorHandler(w, r, err)
		return
	}
	template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		supports.ErrorHandler(w, r, err)
		return
	}
	w.WriteHeader(200)
	return
}
