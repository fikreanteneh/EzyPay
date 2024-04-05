package entities

import "time"

type Auth struct {
	AuthID 	string
	Email       string
	PhoneNumber string
	Password    string
	Verified    bool
	Status      string
	Biometrics  bool
	Twofa       bool
	Oauth 		bool
	Role        string
	CreatedAt   time.Time
}