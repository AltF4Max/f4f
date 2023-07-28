package handlers
import (
  "f4f/models"
  "net/http"
  "encoding/json"
  "gorm.io/gorm"
)
func Register(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  if r.Method == "POST" {
    var user models.Users
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil{
      w.WriteHeader(http.StatusBadRequest)
      answer:=models.Answer{Msg: "Error reading data from request body"}
      json.NewEncoder(w).Encode(answer)
      return
    }
    if len(user.Login)==0 || len(user.PasswordHash)==0 || len(user.Email)==0 || len(user.LastName)==0 || len(user.FirstName)==0{
      w.WriteHeader(http.StatusBadRequest)
      answer:=models.Answer{Msg: "One of the values is empty"}
      json.NewEncoder(w).Encode(answer)
      return
    }
    err=user.RegisterUser(db)
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      answer:=models.Answer{Msg: "Login busy"}//Логин занят или произошла какая то ошибка
      json.NewEncoder(w).Encode(answer)
      return
    }
    answer:=models.Answer{Msg: user.TokenJWT}
    json.NewEncoder(w).Encode(answer)
  } else {//Этого произойти не должно
    w.WriteHeader(http.StatusBadRequest)//400 Запрос клиента содержит синтаксическую ошибку и не может быть выполнен сервером.
    answer:=models.Answer{Msg: "Method not post"}
    json.NewEncoder(w).Encode(answer)
    return
  }
}
}
