package models
import(
  "gorm.io/gorm"
  "time"
)
func (m *Messages) PostMessage(u *Users,db *gorm.DB)error{
  now:=time.Now()
  timeNow := now.Format("2006-01-02 15:04:05")
  m.Login=u.Login
  m.Created=timeNow
  m.Updated=timeNow
  tx := db.Begin()
  if err := tx.Omit("Id").Create(m).Error; err != nil {
      tx.Rollback()// Если возникла ошибка при создании записи, откатываем транзакцию
      return err
  }
  tx.Commit()
  return nil
}
