package controller

// ReqCreate 请求结构体
type ReqCreate struct {
	Name        string  `json:"name" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description" binding:"omitempty"`
	Price       float32 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" binding:"min=1"`
}

// ReqUpdate 修改设备信息
type ReqUpdate struct {
	ID          uint64  `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description" binding:"omitempty"`
	Price       float32 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" binding:"min=1"`
}
