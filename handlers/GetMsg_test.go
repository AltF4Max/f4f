package handlers
import(
  "testing"
  "regexp"
  "github.com/DATA-DOG/go-sqlmock"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "net/http/httptest"
  "net/http"
  "f4f/models"
  "context"

)
func TestGetMsg(t *testing.T) {
  testStruct:= []TestStructMessage{
    {statusCode: 200, contextKey: "User" , fakeDb: 1},
    {statusCode: 404, contextKey: "User"},
    {statusCode: 500},
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

    u := models.Users{
        Id:           23,
        Login:        "s",
        PasswordHash: "s",
        TokenJWT:     "s",
        Email:        "s",
        LastName:     "s",
        FirstName:    "s",
    }
    for _, oneTest := range testStruct {
    ctx := context.WithValue(context.Background(), oneTest.contextKey , u)
  if oneTest.fakeDb == 1 {
    rows := sqlmock.NewRows([]string{"id", "login", "header", "message","created","updated"}).
    AddRow(11,"s", "sd", "s","Model 1", "sd").
    AddRow(12,"s", "sd", "s","Model 1", "sd")

    mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `messages` WHERE login = ?")).
    WithArgs("s").
    WillReturnRows(rows)
}
    req, err := http.NewRequest("GET", "/api/messages/", nil)
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    req = req.WithContext(ctx)

    rr := httptest.NewRecorder()
    handler := GetMsg(gormDB)
    handler.ServeHTTP(rr, req)

    if rr.Code != oneTest.statusCode {
        t.Errorf("Expected status code %v, got %v", oneTest.statusCode, rr.Code)
    }
  }
}
