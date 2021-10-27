package middleware

type HashPasswordFunc func(password string) string

func HashPassword() HashPasswordFunc {
	return func(password string) string {
		return "hashed-password"
	}
}
