package repositories


type AuthRepositoryInterface interface  {
	GenericRepositoryInterface
	FindByEmail(email string) (any, error)
	FindByPhoneNumber(phoneNumber string) (any, error)
}