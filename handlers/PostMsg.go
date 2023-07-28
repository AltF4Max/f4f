package handlers
import (
  "f4f/models"
  "net/http"
  "encoding/json"
  "gorm.io/gorm"
)
func PostMsg(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    user, ok := r.Context().Value("User").(models.Users)
        if !ok {
          w.WriteHeader(http.StatusInternalServerError)
          answer:=models.Answer{Msg: "Error 500 please try again later"}
          json.NewEncoder(w).Encode(answer)
          return
        }
    message:=models.Messages{}
    err := json.NewDecoder(r.Body).Decode(&message)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      answer:=models.Answer{Msg: "Error reading data from request body"}
      json.NewEncoder(w).Encode(answer)
      return
    }
    if len(message.Header)==0 || len(message.Message)==0{
      w.WriteHeader(http.StatusBadRequest)
      answer:=models.Answer{Msg: "Title or message is empty"}
      json.NewEncoder(w).Encode(answer)
      return
    }
    err=message.PostMessage(user.ReturnUsers(), db)
    if err!=nil{
      w.WriteHeader(http.StatusInternalServerError)
      answer:=models.Answer{Msg: "Error 500 please try again later"}
      json.NewEncoder(w).Encode(answer)
      return
    }
    answer:=models.Answer{Msg: "Message sent"}
    json.NewEncoder(w).Encode(answer)
  }
}
