package dbgorm
import (
  "database/sql"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "f4f/config"
)
func ConnectionDB()(*gorm.DB, *sql.DB, error){
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := db.DB()//Стандарт "database/sql"
  if err != nil {
		return nil, nil, err
	}
  return db, sqlDB, nil
}
