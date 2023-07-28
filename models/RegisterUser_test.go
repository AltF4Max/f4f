package models
import(
  "testing"
  "github.com/DATA-DOG/go-sqlmock"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "regexp"
  "fmt"
  "errors"
)
type TestStructReg struct{
  errFinal    error
  fakeDb      int
}
func TestRegisterUser(t *testing.T) {
  testStructReg:= []TestStructReg{
    {fakeDb: 1},
    {errFinal: errors.New("Error"),fakeDb: 2},
    {errFinal: errors.New("Login busy"),fakeDb: 3},
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
  user:= Users{
     Login:        "user123",
     PasswordHash: "s",
     Email:        "s",
     LastName:     "s",
     FirstName:    "s",
 }
for _, oneTest := range testStructReg {
  if oneTest.fakeDb == 1 {
  rows := sqlmock.NewRows([]string{})
  mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE login = ? ORDER BY `users`.`id` LIMIT 1")).
  WithArgs("user123").
  WillReturnRows(rows)

  mock.ExpectBegin()
  mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).
  WillReturnResult(sqlmock.NewResult(1, 1))
  mock.ExpectCommit()
} else if oneTest.fakeDb == 2 {
  rows := sqlmock.NewRows([]string{})
  mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE login = ? ORDER BY `users`.`id` LIMIT 1")).
  WithArgs("user123").
  WillReturnRows(rows)
} else if oneTest.fakeDb == 3{
  rows := sqlmock.NewRows([]string{"id", "login", "password_hash", "token_jwt","email","last_name","first_name"}).
  AddRow(1,"user123", "UCiVPF", "s","Model 1", "sd", "s")
  mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE login = ? ORDER BY `users`.`id` LIMIT 1")).
  WithArgs("user123").
  WillReturnRows(rows)
}
err=user.RegisterUser(gormDB)
if fmt.Sprintf("%v", err) != fmt.Sprintf("%v", oneTest.errFinal){
    t.Errorf("Expected error %v, got %v", oneTest.errFinal, err)
}
}
}
