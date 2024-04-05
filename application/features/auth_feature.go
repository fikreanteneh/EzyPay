package features

import (
	"EzyPay/application/dtos"
	"EzyPay/application/repositories"
)

type AuthFeature struct {
	AuthRepositoryInterface repositories.AuthRepositoryInterface
}

func (feature *AuthFeature) Register(currentUser any, dto *dtos.RegisterDTO, param any) {
	

}

func (feature *AuthFeature) LogInWithGoogle(currentUser any, dto *dtos.LogInWithGoogleDTO, param any) {

}

func (feature *AuthFeature) LogInByEmail(currentUser any, dto *dtos.LogInWithEmailDTO, param any) {

}

func (feature *AuthFeature) LogInByPhoneNumber(currentUser any, dto *dtos.LogInWithPhoneNumberDTO, param any) {

}

func (feature *AuthFeature) RefreshToken(currentUser *dtos.CurrentUserDTO, dto *dtos.RefreshTokenDTO, param any) {

}

func (feature *AuthFeature) ChangePassword(currentUser *dtos.CurrentUserDTO, dto *dtos.ChangePasswordDTO, param any) {

}

func (feature *AuthFeature) ResetPassword(currentUser any, dto any, param any) {

}


func NewAuthFeature(AuthRepositoryInterface *repositories.AuthRepositoryInterface) *AuthFeature {
	return &AuthFeature{
		AuthRepositoryInterface: *AuthRepositoryInterface,
	}

}

