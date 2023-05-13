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

type FyMonitordata struct {
	service.Service
}

// GetPage 获取FyMonitordata列表
func (e *FyMonitordata) GetPage(c *dto.FyMonitordataGetPageReq, p *actions.DataPermission, list *[]models.FyMonitordata, count *int64) error {
	var err error
	var data models.FyMonitordata

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FyMonitordataService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FyMonitordata对象
func (e *FyMonitordata) Get(d *dto.FyMonitordataGetReq, p *actions.DataPermission, model *models.FyMonitordata) error {
	var data models.FyMonitordata

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFyMonitordata error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FyMonitordata对象
func (e *FyMonitordata) Insert(c *dto.FyMonitordataInsertReq) error {
	var err error
	var data models.FyMonitordata
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FyMonitordataService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FyMonitordata对象
func (e *FyMonitordata) Update(c *dto.FyMonitordataUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FyMonitordata{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("FyMonitordataService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FyMonitordata
func (e *FyMonitordata) Remove(d *dto.FyMonitordataDeleteReq, p *actions.DataPermission) error {
	var data models.FyMonitordata

	db := e.Orm.Unscoped().Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFyMonitordata error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
