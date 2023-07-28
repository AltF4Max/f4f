package models
import(
  "gorm.io/gorm"
)
func (m *Messages) DeleteIdMessage(id int, db *gorm.DB) error {
  if err := db.Delete(&m, id).Error; err != nil {
          return err
      }
        return nil
}
