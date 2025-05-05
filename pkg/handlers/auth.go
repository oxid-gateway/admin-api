package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"slices"

	"oxid-gateway-admin-api/pkg/config"
	"oxid-gateway-admin-api/pkg/db"
	"oxid-gateway-admin-api/pkg/dtos"
	"oxid-gateway-admin-api/pkg/services"

	"github.com/go-fuego/fuego"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
)

var jwks jwk.Set

var tokenValidationError = errors.New("Token Validation Error")

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	ctx := context.Background()
	jwksURL := config.GetEnvOrPanic(config.JwksUrl)

	if jwks == nil {
		jwks_response, err := jwk.Fetch(ctx, jwksURL)

		if err != nil {
			slog.Error("failed to fetch JWKS", "error", err)
			return nil, err
		}

		jwks = jwks_response
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			slog.Error("unexpected signing method", "alg", token.Header["alg"])
			return nil, tokenValidationError
		}

		kid := token.Header["kid"]
		key, haveKey := jwks.LookupKeyID(fmt.Sprintf("%v", kid))
		if !haveKey {
			slog.Error("unable to find key")
			return nil, tokenValidationError
		}

		var rawKey any
		err := key.Raw(&rawKey)
		if err != nil {
			slog.Error("unable to extract key", "error", err)
			return nil, tokenValidationError
		}

		return rawKey, nil
	})

	if err != nil {
		slog.Info("failed to validate token", "error", err)
		return nil, tokenValidationError
	}

	return token, nil
}

func ValidateAudience(claims jwt.MapClaims) (bool, error) {
	aud, err := claims.GetAudience()

	if err != nil {
		slog.Info("Error getting audience", "error", err)
		return false, tokenValidationError
	}

	if slices.Contains(aud, config.GetEnvOrPanic(config.TokenAudience)) {
		return true, nil
	}

	return false, nil
}

func valid(authorization string) (*jwt.Token, error) {
	if authorization == "" {
		return nil, nil
	}

	token := strings.TrimPrefix(authorization, "Bearer ")

	parsed_token, err := ValidateJWT(token)

	if err != nil {
		return nil, err
	}

	return parsed_token, nil
}

type RequestContext struct {
	UserEntity *dtos.User
	Roles      []string
}

type RequestContextService struct {
	UsersService *services.UsersService
}

var requestContextServiceInstance *RequestContextService

func NewRequestContextService(r RequestContextService) {
	requestContextServiceInstance = &r
}

func GetRequestContext[B any](c fuego.ContextWithBody[B]) (*RequestContext, error) {
	authHeader := c.Header("Authorization")

	// token, err := valid(authHeader)

	// if err != nil || token == nil {
	// 	w.WriteHeader(401)
	// 	return
	// }

	// claims, ok := token.Claims.(jwt.MapClaims)

	// if !ok || !token.Valid {
	// 	w.WriteHeader(401)
	// 	return
	// }

	// valid_token, err := ValidateAudience(claims)

	// if err != nil {
	// 	w.WriteHeader(500)
	// 	return
	// }

	// if !valid_token {
	// 	w.WriteHeader(401)
	// 	return
	// }

	if authHeader == "" {
		return nil, fuego.UnauthorizedError{Title: "Unauthorized", Detail: "User Unauthorized"}
	}

	username := strings.ReplaceAll(authHeader, "Bearer ", "")

	user, err := requestContextServiceInstance.UsersService.GetUser(username)

	if err != nil {
		slog.Error("Auth Error", "error", err)
		return nil, fuego.InternalServerError{Title: "Auth Error", Detail: "Auth Error"}
	}

	if user == nil {
		// TODO: Get from token claim
		email := username+"@gmail.com"
		name := username+" nomu"

		new_user, err := requestContextServiceInstance.UsersService.CreateUser(db.CreateUserParams{
			Name:     name,
			Username: username,
			Email:    email,
		})

		if err != nil {
			return nil, fuego.InternalServerError{Title: "Auth Error", Detail: "Auth Error"}
		}

		user = &dtos.User{
			ID: new_user.ID,
			Name: new_user.Name,
			Username: new_user.Username,
			Email: new_user.Email,
		}
	}

	return &RequestContext{
		UserEntity: user,
		Roles:      []string{},
	}, nil
}
