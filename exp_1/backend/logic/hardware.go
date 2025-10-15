package logic

import (
	"MES/exp1/dao/mysql"
	"MES/exp1/models"
	"MES/exp1/pkg/snowflake"
	"errors"
	"gorm.io/gorm"
	"time"
)

// CreateHardware 添加设备
func CreateHardware(name, category, description string, price float32, quantity int) error {
	id, err := snowflake.GetID()
	if err != nil {
		return err
	}
	h := models.Product{
		Model: models.Model{
			ID:       id,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		},
		Name:        name,
		Category:    category,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
	return mysql.CreateHardware(h)
}

// UpdateHardware 修改设备信息
func UpdateHardware(id uint64, name, category, description string, price float32, quantity int) error {
	h := models.Product{
		Model: models.Model{
			ID:       id,
			UpdateAt: time.Now(),
		},
		Name:        name,
		Category:    category,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
	return mysql.UpdateHardware(h)
}

func ListHardware(query string) ([]models.Product, int64, error) {
	list, total, err := mysql.ListHardware(query)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.Product{}, 0, nil
	}
	return list, total, err
}

func DeleteHardware(id int64) error {
	return mysql.DeleteHardware(id)
}
