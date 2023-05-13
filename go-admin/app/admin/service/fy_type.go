package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type FyType struct {
	service.Service
}

// GetPage 获取FyType列表
func (e *FyType) GetPage(c *dto.FyTypeGetPageReq, p *actions.DataPermission, list *[]models.FyType, count *int64) error {
	var err error
	var data models.FyType

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FyTypeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FyType对象
func (e *FyType) Get(d *dto.FyTypeGetReq, p *actions.DataPermission, model *models.FyType) error {
	var data models.FyType

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFyType error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FyType对象
func (e *FyType) Insert(c *dto.FyTypeInsertReq) error {
	var err error
	var data models.FyType
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FyTypeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FyType对象
func (e *FyType) Update(c *dto.FyTypeUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FyType{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("FyTypeService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FyType
func (e *FyType) Remove(d *dto.FyTypeDeleteReq, p *actions.DataPermission) error {
	var data models.FyType

	db := e.Orm.Unscoped().Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFyType error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
