package models
import(
  "testing"
)
func TestReturnUsers(t *testing.T) {
 user:= Users{
    Id:           23,
    Login:        "s",
    PasswordHash: "s",
    TokenJWT:     "s",
    Email:        "s",
    LastName:     "s",
    FirstName:    "s",
}
userReturn:=user.ReturnUsers()
if userReturn.Id!=user.Id || userReturn.Login!=user.Login || userReturn.PasswordHash!=user.PasswordHash || userReturn.TokenJWT!=user.TokenJWT || userReturn.Email!=user.Email || userReturn.LastName!=user.LastName || userReturn.FirstName!=user.FirstName{
    t.Errorf("Expected status code %v, got %v", userReturn, user)
}
}
