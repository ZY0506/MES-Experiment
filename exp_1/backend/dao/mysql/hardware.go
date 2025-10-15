package mysql

import (
	"MES/exp1/models"
	"time"
)

// CreateHardware 添加设备
func CreateHardware(h models.Product) error {
	return database.Model(&models.Product{}).Create(&h).Error
}

// UpdateHardware 修改设备信息
func UpdateHardware(h models.Product) error {
	return database.Model(&models.Product{}).Where("id = ?", h.ID).Updates(models.Product{
		Model: models.Model{
			UpdateAt: time.Now(),
		},
		Name:        h.Name,
		Category:    h.Category,
		Description: h.Description,
		Price:       h.Price,
		Quantity:    h.Quantity,
	}).Error
}

// ListHardware 查询设备信息
func ListHardware(query string) ([]models.Product, int64, error) {
	var list []models.Product
	var total int64
	err := database.Model(&models.Product{}).Where("name like ?", "%"+query+"%").Find(&list).Count(&total).Error
	return list, total, err
}

// DeleteHardware 删除设备
func DeleteHardware(id int64) error {
	return database.Model(&models.Product{}).Where("id = ?", id).Delete(&models.Product{}).Error
}
