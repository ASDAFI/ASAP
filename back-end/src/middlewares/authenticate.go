package middlewares

import (
	"farm/configs"
	"farm/src/FarmContext/users"
	"farm/src/infrastructure"
	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

)

func Authenticate(ctx context.Context) (context.Context, error) {
	method, ok := grpc.Method(ctx)
	if ok && method == "/service.farm.Farm/Login" {
		return ctx, nil
	}
	encodedToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	type TokenWithPayload struct {
		users.AuthToken
		FarmID uint `json:"company_id"`
		UserID uint `json:"user_id"`
	}
	tk := &TokenWithPayload{}
	_, err = jwt.ParseWithClaims(encodedToken, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Config.Credential.TokenSecret), nil
	})
	err2 := infrastructure.PostgresDBProvider.DB.Table(users.AuthToken{}.TableName()).Take(tk).Error
	if err != nil || err2 != nil {
		log.Info(err)
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	ctx = context.WithValue(ctx, "username", tk.Username)
	ctx = context.WithValue(ctx, "user_id", tk.UserID)
	ctx = context.WithValue(ctx, "company_id", tk.FarmID)

	return ctx, nil
}
