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
func TestLogin(t *testing.T) {
  testStruct:= []TestStruct{
    {strJson: `{"login": "Modefl1", "password": "admin"}`, statusCode: 200, fakeDb: 1},
    {strJson: `{"login": "Modefl1", "password": "admin"}`, statusCode: 400},
    {strJson: `{"login": "", "password": "admin1"}`, statusCode: 400},
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
  rows := sqlmock.NewRows([]string{"id", "login", "password_hash", "token_jwt","email","last_name","first_name"}).
  AddRow(1,"Modefl1", "$2a$10$s0BO.GcCwDawOOAQwNpoc.I/uKUCiVPFZ0hljtiBvcYtl1lY0km/6", "s","Model 1", "sd", "s")
  mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE login = ? ORDER BY `users`.`id` LIMIT 1")).
  WithArgs("Modefl1").
  WillReturnRows(rows)
}

    req, err := http.NewRequest("POST", "/login", strings.NewReader(oneTest.strJson))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    rr := httptest.NewRecorder()
    handler := Login(gormDB)
    handler.ServeHTTP(rr, req)

    if rr.Code != oneTest.statusCode{
        t.Errorf("Expected status code %v, got %v", oneTest.statusCode, rr.Code)
    }
  }
}
