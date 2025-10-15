package mysql

import (
	"MES/exp1/models"
	"time"
)

// PageResult 分页查询结果
type PageResult struct {
	List  []models.Product `json:"list"`  // 查询结果列表
	Total int64            `json:"total"` // 匹配的总数
	Pages int              `json:"pages"` // 总页数
}

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
func ListHardware(name string, pageNum, pageSize int) (*PageResult, error) {
	//var list []models.Product
	//var total int64
	//err := database.Model(&models.Product{}).Where("name like ?", "%"+name+"%").Find(&list).Count(&total).Error
	//return list, total, err

	result := new(PageResult)

	tx := database.Model(&models.Product{})

	// 模糊查询
	tx = tx.Where("name like ?", "%"+name+"%")

	// 计算总数
	tx = tx.Count(&result.Total)

	// 分页
	offset := (pageNum - 1) * pageSize
	if err := tx.Offset(offset).Limit(pageSize).Find(&result.List).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteHardware 删除设备
func DeleteHardware(id int64) error {
	return database.Model(&models.Product{}).Where("id = ?", id).Delete(&models.Product{}).Error
}
