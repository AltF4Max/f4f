package main
import (
  "encoding/json"
  "f4f/config"
  "context"
  "fmt"
  "f4f/dbgorm"
  "github.com/gorilla/mux"
  "net/http"
  "github.com/dgrijalva/jwt-go"
  "f4f/handlers"
  "f4f/models"
  "strings"
  "gorm.io/gorm"
)
type CustomClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}
func main(){
  fmt.Println("Start")
  db, sqlDB, err:= dbgorm.ConnectionDB()
  if err!=nil{
    fmt.Println("dbgorm.ConnectionDB()", err)
    return
  }
  defer sqlDB.Close()

 r := mux.NewRouter()
 configMy:=config.ConfigMy()
 server := &http.Server{
        Addr:         configMy.Addr,
        Handler:      r,
        ReadTimeout:  configMy.ReadTimeout,
        WriteTimeout: configMy.WriteTimeout,
    }

 r.HandleFunc("/login", handlers.Login(db)).Methods("POST")
 r.HandleFunc("/register", handlers.Register(db)).Methods("POST")
 r.Use(verifyTokenFromDB(db, configMy.JwtSecret, "/login", "/register"))
 r.HandleFunc("/api/messages", handlers.PostMsg(db)).Methods("POST")
 r.HandleFunc("/api/messages", handlers.GetMsg(db)).Methods("GET")
 r.HandleFunc("/api/messages/{id}", handlers.GetMessageId(db)).Methods("GET")
 r.HandleFunc("/api/messages/{id}", handlers.PutMessageId(db)).Methods("PUT")
 r.HandleFunc("/api/messages/{id}", handlers.DeleteMessageId(db)).Methods("DELETE")

    err = server.ListenAndServe()
      if err != nil {
          fmt.Println("Error starting server:", err)
      }
}

func verifyTokenFromDB(db *gorm.DB, jwtSecret []byte, excludePaths ...string) mux.MiddlewareFunc {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

      for _, path := range excludePaths {
        if r.URL.Path == path {
          next.ServeHTTP(w, r)
          return
        }
      }

      tokenString := r.Header.Get("Authorization")
      tokenString = strings.TrimPrefix(tokenString, "Bearer ")
      token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
      })
      if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        answer:=models.Answer{Msg: "Unauthorized"}
        json.NewEncoder(w).Encode(answer)
        return
      }
      if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        u:=models.Users{}
        u.TokenJWT=tokenString
        if err := db.Where("token_jwt = ?", u.TokenJWT).First(&u).Error; err != nil {
          w.WriteHeader(http.StatusUnauthorized)
          answer:=models.Answer{Msg: "Unauthorized"}
          json.NewEncoder(w).Encode(answer)
          return
        }
        if claims.UserID != u.Login {
          w.WriteHeader(http.StatusUnauthorized)
          answer:=models.Answer{Msg: "Unauthorized"}
          json.NewEncoder(w).Encode(answer)
          return
        }
        ctx := context.WithValue(r.Context(), "User", u)
        next.ServeHTTP(w, r.WithContext(ctx))
        return
    }
    w.WriteHeader(http.StatusUnauthorized)
    answer:=models.Answer{Msg: "Unauthorized"}
    json.NewEncoder(w).Encode(answer)
    return
  })
}
}
