package service

import (
	"errors"
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"
	"time"
)

type Claims struct {
	UserID int64 `json:"userId"`
	jwt.StandardClaims
}

/**
 * @DEPRECATED @TODO remove
 */
func (service *Service) isRegisteredAuth(user string, pass string) (*models.UserResponse, error) {
	id, err := service.storage.Login(user, pass)
	if err != nil {
		return nil, err
	}
	email := "email"
	return &models.UserResponse{
		ID:    id,
		Email: &email,
	}, nil
}

func (service *Service) JWTAuth(token string) (*models.UserResponse, error) {
	log.Infof("JWTAuth with token: %s", token)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return service.jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	//if err != nil {
	//	if err == jwt.ErrSignatureInvalid {
	//		w.WriteHeader(http.StatusUnauthorized)
	//		return
	//	}
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	if !tkn.Valid {
		// @TODO Add custom error handling
		// w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("invalid token")
	}
	return &models.UserResponse{
		Email: swag.String("some"),
		ID:    swag.Int64(claims.UserID),
		Token: swag.String(token),
	}, nil
}

func (service *Service) LoginHandler(_ operations.LoginParams, user *models.UserResponse) middleware.Responder {
	payload := models.UserResponse{
		Email: user.Email,
		ID:    user.ID,
	}
	return operations.NewLoginOK().WithPayload(&payload)
}

func (service *Service) SignUpHandler(params operations.SignupParams) middleware.Responder {

	id, err := service.storage.SignUp(*params.User.Email, params.User.Password.String())

	if err != nil {
		payload := models.ErrorResponse{Message: swag.String("Email already taken")}
		return operations.NewSignupDefault(422).WithPayload(&payload)
	}

	token, err := service.createJwtToken(*id)
	if err != nil {
		payload := models.ErrorResponse{Message: swag.String("JWT Token not created")}
		return operations.NewSignupDefault(500).WithPayload(&payload)
	}

	// message := fmt.Sprintf("User [%s] registered with token [%s]", *params.User.Email, token)
	payload := models.UserResponse{
		ID: id,
		Email: params.User.Email,
		Token: token,
	}
	return operations.NewSignupOK().WithPayload(&payload)
}

func (service *Service) createJwtToken(userID int64) (token *string, err error) {
	expirationTime := time.Now().Add(time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := jwtToken.SignedString(service.jwtKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}