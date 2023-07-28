package models
import(
  "testing"
)
func TestReturnLogin(t *testing.T) {
 user:= Users{
    Id:           23,
    Login:        "s",
    PasswordHash: "s",
    TokenJWT:     "s",
    Email:        "s",
    LastName:     "s",
    FirstName:    "s",
}
stringLoginReturn:=user.ReturnLogin()
if stringLoginReturn!=user.Login{
    t.Errorf("Expected status code %v, got %v", stringLoginReturn, user)
}
}
