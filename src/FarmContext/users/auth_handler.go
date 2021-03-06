package users

import (
	"context"
	"errors"
	"farm/configs"
	"farm/src/infrastructure"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthHandler struct {
	userRepo IRepository
}

func NewAuthHandler(userRepo IRepository) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

// todo : separate functions
func (h *AuthHandler) Login(ctx context.Context, username string, password string) (error, string) {
	user, err := h.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return err, ""
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("Username or password not valid. "), ""
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
	tokenString, error1 := token.SignedString([]byte(configs.Config.Credential.TokenSecret))
	if error1 != nil {
		log.Info(error1)
	}
	createToken := infrastructure.PostgresDBProvider.DB.Create(tk)

	if createToken.Error != nil {
		log.Info(createToken.Error)
	}
	return nil, tokenString
}

func (h *AuthHandler) Logout(ctx context.Context, userId uint) error {
	user, err := h.userRepo.FindById(ctx, userId)
	if err != nil {
		return err
	}
	token, err := h.userRepo.FindTokenByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	err = infrastructure.PostgresDBProvider.DB.WithContext(ctx).Delete(token).Error
	if err != nil {
		return err
	}

	/*expiresAt := time.Now().Unix()
	token.ExpiresAt=expiresAt
	createToken := infrastructure.PostgresDBProvider.DB.Updates(token)*/
	return nil
}
