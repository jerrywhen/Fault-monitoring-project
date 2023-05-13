package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	// "github.com/qiniu/go-sdk/v7/reqid"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	cModels "go-admin/common/models"
)

type FyHandle struct {
	service.Service
}

// GetPage 获取FyHandle列表
func (e *FyHandle) GetPage(c *dto.FyHandleGetPageReq, p *actions.DataPermission, list *[]models.FyHandle, count *int64) error {
	var err error
	var data models.FyHandle

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
			cModels.OrderDesc("id"),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FyHandleService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FyHandle对象
func (e *FyHandle) Get(d *dto.FyHandleGetReq, p *actions.DataPermission, model *models.FyHandle) error {
	var data models.FyHandle

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFyHandle error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FyHandle对象
func (e *FyHandle) Insert(c *dto.FyHandleInsertReq) error {
	var err error
	var data models.FyHandle
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FyHandleService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FyHandle对象
func (e *FyHandle) Update(c *dto.FyHandleUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FyHandle{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("FyHandleService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FyHandle
func (e *FyHandle) Remove(d *dto.FyHandleDeleteReq, p *actions.DataPermission) error {
	var data models.FyHandle

	db := e.Orm.Unscoped().Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFyHandle error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
