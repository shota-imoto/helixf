package helixf_user

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/shota-imoto/helixf/lib/utils/line"
)

type User struct {
	Id     uint   `gorm:"primaryKey" sql:"type:uint"`
	Name   string `sql:"type:string"`
	LineId string `sql:"type:string" gorm:"unique"`
}

type ParseInterface interface {
	parseToken() (*jwt.Token, error)
}

type ParseStruct struct {
	Parser ParseInterface
}

// 本番用
type BasicParser struct {
	IdToken string
}

func (p *BasicParser) parseToken() (*jwt.Token, error) {
	return line.ParseIdToken(p.IdToken)
}

// テスト用 //
type DummyParser struct {
	Name   string
	LineId string
}

func (p *DummyParser) parseToken() (*jwt.Token, error) {
	return &jwt.Token{Claims: jwt.MapClaims{"name": p.Name, "sub": p.LineId}}, nil
}

/////////////

func (p *ParseStruct) GetJwtClaims() (jwt.MapClaims, error) {
	token, err := p.Parser.parseToken()
	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, errors.New("failed get claims")
	}

	return claims, err
}

func (p *ParseStruct) BuildUserByIdToken() (User, error) {
	token, err := p.Parser.parseToken()
	if err != nil {
		return User{}, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	user := User{Name: string(claims["name"].(string)), LineId: string(claims["sub"].(string))}

	return user, nil
}
