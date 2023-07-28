package dbgorm
import (
  "database/sql"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)
func ConnectionDB()(*gorm.DB, *sql.DB, error){
  dsn := "tester:secret@tcp(tesst_db:3306)/tesst"//"tester:secret@tcp(tesst_db:3306)/tesst""root:@tcp(127.0.0.1:3306)/something"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := db.DB()//Стандарт "database/sql"
  if err != nil {
		return nil, nil, err
	}
  return db, sqlDB, nil
}
