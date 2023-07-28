package main

import (
    "regexp"
    "net/http"
    "net/http/httptest"
    "github.com/DATA-DOG/go-sqlmock"
    "gorm.io/driver/mysql"
    "testing"
    "github.com/stretchr/testify/assert"
    "gorm.io/gorm"
)
type TestStruct struct{
  Token string
  Url string
  statusCode int
  fakeDb int
}
func TestVerifyTokenFromDB(t *testing.T) {
  testStruct:= []TestStruct{
    {Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJhZG1pbiIsImV4cCI6MTcxMTQ3MjQ4OH0.UI79VMNx3uKlqhB3U_VKihm3cmOLDEy0Nrvh8wELmh8", Url: "/test", statusCode: 200, fakeDb: 1},
    {Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJhZG1kaW4iLCJleHAiOjE3MTE1NjU1MDh9.bz1nRPLyrEklKek-L4Pmd5HG_BR7a48phYSvkN1_j4s", Url: "/test", statusCode: 401,  fakeDb: 1},
    {Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJhZG1pbiIsImV4cCI6MTcxMTQ3MjQ4OH0.UI79VMNx3uKlqhB3U_VKihm3cmOLDEy0Nrvh8wELmh8", Url: "/test", statusCode: 401},
    {Token: "", Url: "/test", statusCode: 401},
    {Url: "/login", statusCode: 200},
  }

    mockDB, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("failed to create mock database: %v", err)
    }
    defer mockDB.Close()

    gormDB, err := gorm.Open(mysql.New(mysql.Config{
        Conn:                      mockDB,
        SkipInitializeWithVersion: true,
    }), &gorm.Config{})
    if err != nil {
        t.Fatalf("failed to open GORM database: %v", err)
    }

    testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })

    excludePaths := []string{"/login"} // Путь, который мы хотим исключить из проверки токена
    handler := verifyTokenFromDB(gormDB, []byte("f4keraven"), excludePaths...)(testHandler)

    for _, oneTest := range testStruct {
      if oneTest.fakeDb == 1 {
      rows := sqlmock.NewRows([]string{"id", "login", "password_hash", "token_jwt","email","last_name","first_name"}).
      AddRow(1,"admin", "f", "s4fr","Model 1", "sd", "s")
      mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE token_jwt = ? ORDER BY `users`.`id` LIMIT 1")).
      WithArgs(oneTest.Token).
      WillReturnRows(rows)
        }
    req, _ := http.NewRequest("GET", oneTest.Url, nil)
    req.Header.Set("Authorization", "Bearer "+oneTest.Token)
    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)
    assert.Equal(t, oneTest.statusCode, rr.Code, "Error")
  }
}
