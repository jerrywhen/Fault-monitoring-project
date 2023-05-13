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

type FyFabric struct {
	service.Service
}

// GetPage 获取FyFabric列表
func (e *FyFabric) GetPage(c *dto.FyFabricGetPageReq, p *actions.DataPermission, list *[]models.FyFabric, count *int64) error {
	var err error
	var data models.FyFabric

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FyFabricService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FyFabric对象
func (e *FyFabric) Get(d *dto.FyFabricGetReq, p *actions.DataPermission, model *models.FyFabric) error {
	var data models.FyFabric

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFyFabric error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FyFabric对象
func (e *FyFabric) Insert(c *dto.FyFabricInsertReq) error {
	var err error
	var data models.FyFabric
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FyFabricService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FyFabric对象
func (e *FyFabric) Update(c *dto.FyFabricUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FyFabric{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("FyFabricService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FyFabric
func (e *FyFabric) Remove(d *dto.FyFabricDeleteReq, p *actions.DataPermission) error {
	var data models.FyFabric

	db := e.Orm.Unscoped().Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFyFabric error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
