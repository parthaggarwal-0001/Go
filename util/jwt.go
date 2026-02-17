package util
import (
	"jwt"
)

func jwt(){
	var secretKey = []byte("mysecretkey")

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	return token.SignedString(secretKey)
}

}