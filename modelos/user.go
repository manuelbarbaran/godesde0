package modelos

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}

func (usuario *User) NewUser(id int, username string, email string, password string, createdAt time.Time, status string) {

	usuario.ID = id
	usuario.Username = username
	usuario.Email = email
	usuario.Password = password
	usuario.CreatedAt = createdAt
	usuario.Status = status
}
