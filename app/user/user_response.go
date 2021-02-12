package user

type ResponseUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserResponseFormatter(user User) ResponseUser {
	formatter := ResponseUser{
		Name:  user.Name,
		Email: user.Email,
	}

	return formatter
}
