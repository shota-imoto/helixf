package line_service

import (
	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/helixf_user"
)

func FindUserByIdToken(parser *helixf_user.ParseStruct) (helixf_user.User, error) {
	user := helixf_user.User{}
	claims, err := parser.GetJwtClaims()

	if err != nil {
		return user, nil
	}

	db.Db.Where("line_id = ?", claims["sub"].(string)).First(&user)

	return user, nil
}

// ex
// parser := &helixf_user.ParseStruct{Parser: &helixf_user.BasicParser{IdToken: tokenResponse.IdToken}}
// user, err := line_service.FindOrCreateUserByIdToken("test", parser)

func FindOrCreateUserByIdToken(parser *helixf_user.ParseStruct) (helixf_user.User, error) {
	user, err := FindUserByIdToken(parser)

	if err != nil {
		return helixf_user.User{}, err
	}

	if user.Id == 0 {

		user, err := parser.BuildUserByIdToken()

		if err != nil {
			return helixf_user.User{}, err
		}

		db.Db.Create(&user)
	}

	return user, nil
}
