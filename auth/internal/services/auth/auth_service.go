package auth

import (
	"context"
	"fmt"
	au "github.com/Bubotka/Microservices/auth/pkg/go/auth"
	"github.com/Bubotka/Microservices/proxy/pkg/clients/user/grpc"
	"github.com/Bubotka/Microservices/user/domain/models"
	"github.com/go-chi/jwtauth"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	user grpc.UserProviderer
	au.UnimplementedAuthServer
}

func NewAuthService(user grpc.UserProviderer) *AuthService {
	return &AuthService{user: user}
}

func (a *AuthService) Register(ctx context.Context, request *au.UserAuthRequest) (*empty.Empty, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return &empty.Empty{}, err
	}

	err = a.user.CheckUser(models.User{
		Username: request.Username,
		Password: request.Password,
	})

	if err != nil {
		return nil, err
	}

	err = a.user.Create(models.User{
		Username: request.Username,
		Password: string(password),
		IsDelete: false,
	})
	return &empty.Empty{}, err
}

func (a *AuthService) Login(ctx context.Context, request *au.UserAuthRequest) (*au.LoginResponse, error) {
	user, err := a.user.Profile(request.Username)
	fmt.Println("Auth login", err)
	if err != nil {
		return nil, err
	}
	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(request.Password))
	if err != nil {
		return &au.LoginResponse{}, err
	}

	tokenAuth := jwtauth.New("HS256", []byte("mysecretkey"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"username": request.Username})
	return &au.LoginResponse{Token: tokenString}, nil
}
