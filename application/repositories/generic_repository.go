package repositories



type GenericRepositoryInterface interface {
	GetAll() ([]any, error)
	GetById(id string) (any, error)
	Create(data any) (any, error)
	Update(data any) (any, error)
	Delete(id string) (any, error)
}