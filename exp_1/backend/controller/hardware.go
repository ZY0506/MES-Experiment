package controller

import (
	"MES/exp1/logic"
	"MES/exp1/models"
	"MES/exp1/pkg/validate"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

// QueryResult 查询结果
type QueryResult struct {
	List  []models.Product `json:"list"`
	Total int64            `json:"total"`
}

// CreateHardwareView 添加设备
func CreateHardwareView(c *gin.Context) {
	// 参数校验
	var req ReqCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			FailWithMessage("参数错误", c)
			return
		}
		FailWithMessage(validate.RemoveTopStruct(errs.Translate(validate.Trans)), c)
		return
	}
	// 添加设备
	if err := logic.CreateHardware(req.Name, req.Category, req.Description, req.Price, req.Quantity); err != nil {
		zap.L().Error("添加设备失败", zap.Error(err))
		FailWithMessage("添加设备失败", c)
		return
	}
	// 返回响应
	SuccessWithMessage("添加设备成功", nil, c)
}

// UpdateHardwareView 修改设备信息
func UpdateHardwareView(c *gin.Context) {
	// 参数校验
	var req ReqUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			FailWithMessage("参数错误", c)
			return
		}
		FailWithMessage(validate.RemoveTopStruct(errs.Translate(validate.Trans)), c)
		return
	}
	// 修改设备信息
	if err := logic.UpdateHardware(req.ID, req.Name, req.Category, req.Description, req.Price, req.Quantity); err != nil {
		zap.L().Error("修改设备信息失败", zap.Error(err))
		FailWithMessage("修改设备信息失败", c)
		return
	}
	// 返回响应
	SuccessWithMessage("修改设备信息成功", nil, c)
}

// ListHardwareView 查询设备信息
func ListHardwareView(c *gin.Context) {
	// 参数绑定
	var query string
	query = c.Query("name")
	// 查询设备信息
	list, total, err := logic.ListHardware(query)
	if err != nil {
		zap.L().Error("查询设备信息失败", zap.Error(err))
		FailWithMessage("查询设备信息失败", c)
		return
	}
	res := QueryResult{
		List:  list,
		Total: total,
	}
	SuccessWithMessage("查询设备信息成功", res, c)
}

// DeleteHardwareView 删除设备
func DeleteHardwareView(c *gin.Context) {
	// 参数绑定
	var query string
	query = c.Query("id")
	id, err := strconv.Atoi(query)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		FailWithMessage("参数错误", c)
		return
	}
	// 删除设备
	if err := logic.DeleteHardware(int64(id)); err != nil {
		zap.L().Error("删除设备失败", zap.Error(err))
		FailWithMessage("删除设备失败", c)
		return
	}
	SuccessWithMessage("删除设备成功", nil, c)
}
