package handlers
import(
  "testing"
  "regexp"
  "github.com/DATA-DOG/go-sqlmock"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "fmt"
)
type AuthorizationStruct struct{
  Login       string
  Password    string
  TokenAnswer string
  fakeDb      int
}
func TestLoginAuthorization(t *testing.T) {
  authorizationStruct:= []AuthorizationStruct{
    {Login: "Modefl1", Password: "admin", TokenAnswer: "s4fr", fakeDb: 1},
    {Login: "Modefl1", Password: "admin1", TokenAnswer: "", fakeDb: 1},
    {Login: "Modefl1", Password: "admin", TokenAnswer: "", fakeDb: 2},
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

for _, oneTest := range authorizationStruct {
  if oneTest.fakeDb==1{
  rows := sqlmock.NewRows([]string{"id", "login", "password_hash", "token_jwt","email","last_name","first_name"}).
  AddRow(1,"Modefl1", "$2a$10$s0BO.GcCwDawOOAQwNpoc.I/uKUCiVPFZ0hljtiBvcYtl1lY0km/6", "s4fr","Model 1", "sd", "s")
  mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE login = ? ORDER BY `users`.`id` LIMIT 1")).
  WithArgs("Modefl1").
  WillReturnRows(rows)
    }
token, err:= authorization(gormDB, oneTest.Login, oneTest.Password)
if token!=oneTest.TokenAnswer{
  fmt.Println(err)
  t.Errorf("Expected token %v, got %v", oneTest.TokenAnswer, token)
}
}
}
