package dtos

import "time"

type Auth struct {
	AuthID      string  `json:"authID"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phoneNumber"`
	Password    string  `json:"password"`
	Verified    bool    `json:"verified"`
	Status      string  `json:"status"`
	Biometrics  bool    `json:"biometrics"`
	Twofa       bool     `json:"twofa"`
	Oauth       bool     `json:"oauth"`
	Role        string   `json:"role"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CurrentUserDTO struct {
	AuthID 	string `json:"authID"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"role"`
	Verified   bool   `json:"verified"`
	Status     string `json:"status"`
}


type RegisterDTO struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type LogInWithPhoneNumberDTO struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type LogInWithEmailDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInWithGoogleDTO struct {
	Token string `json:"token"`
}

type ChangePasswordDTO struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:newPassword`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refreshToken"`
}




