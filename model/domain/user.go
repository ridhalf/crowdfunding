package domain

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Occupation     string `json:"occupation"`
	Email          string `json:"email"`
	PasswordHash   string `json:"password_hash"`
	AvatarFileName string `json:"avatar_file_name"`
	Role           string `json:"role"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Token          string `json:"token"`
}
