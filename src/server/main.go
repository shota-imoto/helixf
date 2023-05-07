package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shota-imoto/helixf/src/server/handlers"
	"github.com/shota-imoto/helixf/src/server/middleware"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/regular_schedule_template", handlers.PostRegularScheduleTemplateHandler).Methods(http.MethodPost)
	r.HandleFunc("/regular_schedule_templates/{id}", handlers.DeleteRegularScheduleTemplateHandler).Methods(http.MethodDelete)
	r.HandleFunc("/groups/register", handlers.RegisterGroups).Methods(http.MethodPost)
	r.HandleFunc("/groups", handlers.GetListGroups).Methods(http.MethodGet)
	r.HandleFunc("/groups/{id}", handlers.GetGroup).Methods(http.MethodGet)
	r.HandleFunc("/groups/{id}/regular_schedule_templates", handlers.GetListRegularScheduleTemplates).Methods(http.MethodGet)
	r.HandleFunc("/callback", handlers.LineCallbackHandler)
	r.HandleFunc("/authenticate", handlers.LineAuthenticationHandler)  // AuthorizatonCode取得
	r.HandleFunc("/assert_auth", handlers.AssertAuthenticationHandler) // AuthorizationCode検証＆AuthorizationToken取得
	r.HandleFunc("/refresh_auth", handlers.RefreshAuthenticationHandler).Methods((http.MethodPost))

	// CORSのpreflightリクエストの受諾が必要なパスはこちらに追加すること
	r.HandleFunc("/regular_schedule_template", handlers.CorsHandler).Methods(http.MethodOptions)
	r.HandleFunc("/regular_schedule_templates/{id}", handlers.CorsHandler).Methods(http.MethodOptions)
	r.HandleFunc("/groups/register", handlers.CorsHandler).Methods(http.MethodOptions)
	r.HandleFunc("/groups", handlers.CorsHandler).Methods(http.MethodOptions)
	r.HandleFunc("/groups/{id}", handlers.CorsHandler).Methods(http.MethodOptions)
	r.HandleFunc("/groups/{id}/regular_schedule_templates", handlers.GetListRegularScheduleTemplates).Methods(http.MethodOptions)
	r.HandleFunc("/refresh_auth", handlers.CorsHandler).Methods(http.MethodOptions)

	// r.Use(mux.CORSMethodMiddleware(r))
	r.Use(middleware.SetCorsHandler)
	r.Use(middleware.GetAuthUser)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(srv.ListenAndServe())
}
