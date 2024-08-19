package auth

type JwtService interface {
	GenerateToken(userID int) (string, error)
}
