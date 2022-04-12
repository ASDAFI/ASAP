package models

import (
	"farm/configs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetUser(username string) (*User, error) {
	user := &User{}
	result := PostgresDBProvider.DB.Where("username = ?", username).First(user)
	return user, result.Error
}

func Login(user *User, password string) (status bool, errMessage string, tok string) {

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, "Username or password not valid.", ""
	}

	tk := &AuthToken{
		Username: user.Username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Id:        fmt.Sprintf("%s", uuid.NewV4()),
		},
	}
	type TokenWithPayload struct {
		*AuthToken
		CompanyId uint `json:"company_id"`
		UserId    uint `json:"user_id"`
	}
	tokenWithPayload := TokenWithPayload{
		CompanyId: user.FarmID,
		UserId:    user.ID,
	}
	tokenWithPayload.AuthToken = tk
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenWithPayload)
	tokenString, error := token.SignedString([]byte(configs.Config.Credential.TokenSecret))
	if error != nil {
		log.Info(error)
	}
	createToken := PostgresDBProvider.DB.Create(tk)

	if createToken.Error != nil {
		log.Info(createToken.Error)
	}
	return true, "", tokenString
}
