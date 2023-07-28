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
  "f4f/models"
  "context"
)
func TestPostMsg(t *testing.T) {
  testStruct:= []TestStructMessage{
    {strJson: `{"Header": "ddd", "Message": "testpass"}`, statusCode: 200, contextKey: "User" , fakeDb: 1},
    {strJson: `{"Header": "ddd", "Message": "testpass"}`, statusCode: 500, contextKey: "User"},
    {strJson: `{"Header": "ddd", "Message": ""}`, statusCode: 400, contextKey: "User"},
    {strJson: ``, statusCode: 400, contextKey: "User"},
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
    mock.ExpectBegin()
    mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `messages`")).
    WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
}
    req, err := http.NewRequest("POST", "/api/messages/", strings.NewReader(oneTest.strJson))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    req = req.WithContext(ctx)

    rr := httptest.NewRecorder()
    handler := PostMsg(gormDB)
    handler.ServeHTTP(rr, req)

    if rr.Code != oneTest.statusCode {
        t.Errorf("Expected status code %v, got %v", oneTest.statusCode, rr.Code)
    }
  }
}
