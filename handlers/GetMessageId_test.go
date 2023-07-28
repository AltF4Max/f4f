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
  "github.com/gorilla/mux"
  "fmt"
)
func TestGetMessageId(t *testing.T) {
  testStruct:= []TestStructMessage{
    {statusCode: 200, urlId: "11", contextKey: "User" , fakeDb: 1},
    {statusCode: 404, urlId: "11", contextKey: "User" , fakeDb: 2},
    {statusCode: 403, urlId: "11", contextKey: "User"},
    {statusCode: 400, urlId: "11a", contextKey: "User"},
    {statusCode: 500, urlId: "11"},
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
    AddRow(11,"s", "sd", "s","Model 1", "sd")
    mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `messages` WHERE id = ? ORDER BY `messages`.`id` LIMIT 1")).
    WithArgs(11).
    WillReturnRows(rows)

    rows2 := sqlmock.NewRows([]string{"id", "login", "header", "message","created","updated"}).
    AddRow(11,"s", "sd", "s","Model 1", "sd")
    mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `messages` WHERE id = ? ORDER BY `messages`.`id` LIMIT 1")).
    WithArgs(11).
    WillReturnRows(rows2)

} else if oneTest.fakeDb == 2{
  rows := sqlmock.NewRows([]string{"id", "login", "header", "message","created","updated"}).
  AddRow(11,"s", "sd", "s","Model 1", "sd")
  mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `messages` WHERE id = ? ORDER BY `messages`.`id` LIMIT 1")).
  WithArgs(11).
  WillReturnRows(rows)
}

    url := fmt.Sprintf("/api/messages/"+oneTest.urlId)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }
    vars := map[string]string{
            "id": oneTest.urlId,
        }

    req = req.WithContext(ctx)
    req = mux.SetURLVars(req, vars)

    rr := httptest.NewRecorder()
    handler := GetMessageId(gormDB)
    handler.ServeHTTP(rr, req)

    if rr.Code != oneTest.statusCode {
        t.Errorf("Expected status code %v, got %v", oneTest.statusCode, rr.Code)
    }
  }
}
