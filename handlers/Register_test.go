package handlers
import(
  "testing"
  "regexp"
  "github.com/DATA-DOG/go-sqlmock"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "net/http/httptest"
  "net/http"
  "strings"
)
type TestStruct struct{
  strJson    string
  statusCode int
  fakeDb     int
}
func TestRegisterHandler(t *testing.T) {
  testStruct:= []TestStruct{
    {strJson: `{"login": "Modefl1", "password": "testpass", "email": "test@example.com", "lastname": "Test", "firstname": "User"}`, statusCode: 200, fakeDb: 1},
    {strJson: `{"login": "Modefl1", "password": "testpass", "email": "test@example.com", "lastname": "Test", "firstname": "User"}`, statusCode: 500},
    {strJson: `{"login": "", "password": "testpass", "email": "test@example.com", "lastname": "Test", "firstname": "User"}`, statusCode: 400},
    {strJson: ``, statusCode: 400},
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

    for _, oneTest := range testStruct {
if oneTest.fakeDb==1{
  rows := sqlmock.NewRows([]string{})
  mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE login = ? ORDER BY `users`.`id` LIMIT 1")).
  WithArgs("Modefl1").
  WillReturnRows(rows)

  mock.ExpectBegin()
  mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).
  WillReturnResult(sqlmock.NewResult(1, 1))
  mock.ExpectCommit()
}
/*rows := sqlmock.NewRows([]string{"id", "login", "password_hash", "token_jwt","email","last_name","first_name"}).
  AddRow(1,"Modefl1", "sd", "s","Model 1", "sd", "s")
  mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE login = ? ORDER BY `users`.`id` LIMIT 1")).
  WithArgs("Modefl1").
  WillReturnRows(rows)*/
    req, err := http.NewRequest("POST", "/register", strings.NewReader(oneTest.strJson))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    rr := httptest.NewRecorder()
    handler := Register(gormDB)
    handler.ServeHTTP(rr, req)
    //fmt.Println(rr.Body.String())
    if rr.Code != oneTest.statusCode{
        t.Errorf("Expected status code %v, got %v", oneTest.statusCode, rr.Code)
    }
  }
}
