package model

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username" validate:"required"`
	NameFirst string    `json:"nameFirst" validate:"required"`
	NameLast  string    `json:"nameLast" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserViewModel struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	NameFirst string `json:"nameFirst"`
	NameLast  string `json:"nameLast"`
	NameFull  string `json:"nameFull"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

func UserToViewModel(user User) UserViewModel {
	return UserViewModel{
		Id:        user.Id,
		Username:  user.Username,
		NameFirst: user.NameFirst,
		NameLast:  user.NameLast,
		NameFull:  user.NameFirst + " " + user.NameLast,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}
