package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/service/line_service"
	"github.com/shota-imoto/helixf/src/server/supports"
)

var AuthorizationUserKey string = "user"

func GetAuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parser := &helixf_user.ParseStruct{Parser: &helixf_user.BasicParser{IdToken: r.Header.Get("Authorization")}}
		user, err := line_service.FindOrCreateUserByIdToken(parser)

		if err != nil {
			if err.Error() == "Token is expired" {
				supports.UnauthorizedHandler(w, r, err)
				return
			} else {
				fmt.Println(err)
			}
		}

		ctxWithUser := context.WithValue(r.Context(), AuthorizationUserKey, user)
		next.ServeHTTP(w, r.WithContext(ctxWithUser))
	})
}
