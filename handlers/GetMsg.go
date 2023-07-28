package handlers
import (
  "f4f/models"
  "net/http"
  "encoding/json"
  "gorm.io/gorm"
)
func GetMsg(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    user, ok := r.Context().Value("User").(models.Users)
        if !ok {
          w.WriteHeader(http.StatusInternalServerError)
          answer:=models.Answer{Msg: "Error 500 please try again later"}
          json.NewEncoder(w).Encode(answer)
          return
        }
    result, err:=models.GetMessage(user.ReturnLogin(),db)
    if err!=nil{
      w.WriteHeader(http.StatusNotFound)
      answer:=models.Answer{Msg: "The requested resource was not found on the server"}
      json.NewEncoder(w).Encode(answer)
      return
    }
    json.NewEncoder(w).Encode(result)
}
}
