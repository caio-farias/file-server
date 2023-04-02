package repository

type UserReader interface {
	FindByEmail(email string) (*interface{}, error)
	FindByChangePasswordHash(hash string) (*interface{}, error)
	FindByValidationHash(hash string) (*interface{}, error)
}

type UserRepository interface {
	Reader
	UserReader
	Writer
}
