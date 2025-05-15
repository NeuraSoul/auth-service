package services

import (
	"context"
	"log"

	pd "auth-service/pkg/grpc"
	"auth-service/internal/domain/models"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*models.Users,error)
}


type AuthService struct {
	rep AuthRepository
	pd.UnimplementedAuthServiceServer
}

func NewAuthServer(rep AuthRepository) *AuthService {
	return &AuthService{rep: rep}
}

func (s *AuthService) LoginUser (ctx context.Context,req *pd.UserRequest) (*pd.UserResponse,error) {
	user,err := s.rep.GetUserByEmail(req.Email)
	if err != nil {
		log.Println("fait to get user by email:",err)
	}
	if user.Password != req.Password {
		log.Printf("missing passsword : %v",err)
	}


	return &pd.UserResponse{Result:false,Token:"hello"},nil

}