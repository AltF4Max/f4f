package models
import(
  "gorm.io/gorm"
  "errors"
)
func (u *Users) ChekReqId(id int,db *gorm.DB) error{
  var message Messages
  if err := db.First(&message, "id = ?", id).Error; err != nil {
return err
}
if u.Login!=message.Login{
  err := errors.New("Произошла ошибка")
  return err
}
return nil
}
