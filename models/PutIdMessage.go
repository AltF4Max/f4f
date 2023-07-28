package models
import(
  "time"
  "gorm.io/gorm"
)
func (m *Messages) PutIdMessage(id int, db *gorm.DB) error {
  header:=m.Header
  message:=m.Message
  if err := db.First(&m, id).Error; err != nil {
    return err
  }
  now:=time.Now()
  timeNow := now.Format("2006-01-02 15:04:05")
m.Updated=timeNow
m.Message=message
m.Header=header
  if err := db.Save(&m).Error; err != nil {
    return err
  }
  return nil
}
/*if err := db.Model(&m).Where("id = ?", id).Updates(map[string]interface{}{
    "header": m.Header,
    "message": m.Message,
    "updated": timeNow,
}).Error; err != nil {
    return err
}*/
