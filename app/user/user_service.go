package user

type Services interface {
	CreateUser(req RequestUser) User
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) CreateUser(req RequestUser) User {
	user := User{}
	user.Name = req.Name
	user.Email = req.Email
	user.Password = req.Password

	newUser := s.repository.InsertUser(user)

	return newUser
}
