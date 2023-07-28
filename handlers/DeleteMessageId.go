package handlers
import (
  "f4f/models"
  "net/http"
  "github.com/gorilla/mux"
  "strconv"
  "encoding/json"
  "gorm.io/gorm"
)
func DeleteMessageId(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    user, ok := r.Context().Value("User").(models.Users)
        if !ok {
          w.WriteHeader(http.StatusInternalServerError)
          answer:=models.Answer{Msg: "Error 500 please try again later"}
          json.NewEncoder(w).Encode(answer)
          return
        }
        id := mux.Vars(r)["id"]
        i, err := strconv.Atoi(id)
      if err!=nil{
        w.WriteHeader(http.StatusBadRequest)
        answer:=models.Answer{Msg: "Invalid ID"}
        json.NewEncoder(w).Encode(answer)
        return
      }
      err=user.ChekReqId(i, db)
      if err!=nil{
        w.WriteHeader(http.StatusForbidden)
        answer:=models.Answer{Msg: "No access"}
        json.NewEncoder(w).Encode(answer)
        return
      }
      message:=models.Messages{}
      err=message.DeleteIdMessage(i, db)
      if err!=nil{
        w.WriteHeader(http.StatusInternalServerError)
        answer:=models.Answer{Msg: "Repeat later"}
        json.NewEncoder(w).Encode(answer)
        return
      }
      answer:=models.Answer{Msg: "Deleted"}
      json.NewEncoder(w).Encode(answer)
}
}
