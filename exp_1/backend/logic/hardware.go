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

func ListHardware(name string, pageNum, pageSize int) (*mysql.PageResult, error) {
	// 校验并设置默认值
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	result, err := mysql.ListHardware(name, pageNum, pageSize)
	if err != nil {
		// 如果是记录未找到错误，返回空结果
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &mysql.PageResult{
				List:  []models.Product{},
				Total: 0,
			}, nil
		}
		return nil, err
	}
	return result, err
}

func DeleteHardware(id int64) error {
	return mysql.DeleteHardware(id)
}
