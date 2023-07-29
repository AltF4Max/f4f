package config
import (
  "time"
)
var Dsn string = "tester:secret@tcp(tesst_db:3306)/tesst"//"tester:secret@tcp(tesst_db:3306)/tesst""root:@tcp(127.0.0.1:3306)/something"
var jwtSecret = []byte("f4keraven")
type ServerConfig struct {
    Addr         string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
    JwtSecret    []byte
}
func ConfigMy()ServerConfig{
  config := ServerConfig{
        Addr:         ":8080",
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        JwtSecret:   jwtSecret,
    }
    return config
}
func ReturnJwtSecret()[]byte{
    return jwtSecret
}
