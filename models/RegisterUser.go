package models
import (
"f4f/config"
"errors"
"time"
"github.com/dgrijalva/jwt-go"
"golang.org/x/crypto/bcrypt"
"gorm.io/gorm"
)
func (u *Users) RegisterUser(db *gorm.DB) error{
ok:=busyUser(u.Login, db)

if ok==true{
  return errors.New("Login busy")
}
var PasswordByte []byte
PasswordByte, err:=CreationPws(u.PasswordHash)
if err!=nil{
  return err
}
u.PasswordHash=string(PasswordByte)
u.TokenJWT, err= CreationToken(u.Login)
if err!=nil{
  return err
}
if err := db.Create(&u).Error; err != nil {
  return errors.New("Error")
}
return nil
}
func busyUser(u string, db *gorm.DB)bool{
  var user Users
  result := db.Where("login = ?", u).First(&user)
  if result.Error == nil {
    return true
  }
  return false
}
func CreationPws(p string) ([]byte, error){
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
  if err != nil {
    err = errors.New("Ошибка при хешировании пароля")
      return nil,err
  }
  return hashedPassword, nil
}
func CreationToken(u string) (string, error){
  claims := &CustomClaims{
    UserID: u,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * 6000).Unix(),
  },
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, err := token.SignedString(config.ReturnJwtSecret())
if err != nil {
  err = errors.New("Ошибка при создании токена")
return "", err
}
return tokenString, nil
}
