package user

import "time"

type ResponseUser struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	AuthToken string     `json:"auth_token"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func UserResponseFormatter(user User, auth_token string) ResponseUser {
	formatter := ResponseUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		AuthToken: auth_token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	return formatter
}
