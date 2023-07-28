package models
import(
  "gorm.io/gorm"
)
func (m *Messages) GetIdMessage(id int, db *gorm.DB) error {
    if err := db.First(&m, "id = ?", id).Error; err != nil {
  return err
}
        return nil
}
