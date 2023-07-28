package models
import(
	"github.com/dgrijalva/jwt-go"
)
type Users struct {
	Id           int
	Login        string `json:"login"`
  PasswordHash string `json:"password"`
	TokenJWT     string
	Email        string `json:"email"`
  LastName     string `json:"lastname"`
  FirstName    string `json:"firstname"`
}
type Messages struct {
	Id      int
	Login   string
	Header  string `json:"header"`
	Message string `json:"message"`
	Created string
	Updated string
}
type CustomClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}
type Answer struct {
	Msg string `json:"msg"`
}
