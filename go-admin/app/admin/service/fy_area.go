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

type FyArea struct {
	service.Service
}

// GetPage 获取FyArea列表
func (e *FyArea) GetPage(c *dto.FyAreaGetPageReq, p *actions.DataPermission, list *[]models.FyArea, count *int64) error {
	var err error
	var data models.FyArea

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FyAreaService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FyArea对象
func (e *FyArea) Get(d *dto.FyAreaGetReq, p *actions.DataPermission, model *models.FyArea) error {
	var data models.FyArea

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFyArea error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FyArea对象
func (e *FyArea) Insert(c *dto.FyAreaInsertReq) error {
	var err error
	var data models.FyArea
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FyAreaService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FyArea对象
func (e *FyArea) Update(c *dto.FyAreaUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FyArea{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("FyAreaService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FyArea
func (e *FyArea) Remove(d *dto.FyAreaDeleteReq, p *actions.DataPermission) error {
	var data models.FyArea

	db := e.Orm.Unscoped().Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFyArea error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// 本文中的代码是 Go 语言的服务层，用于处理 FyArea 数据库表的增删改查操作。
// FyArea 结构体继承了 go-admin-core/sdk/service 包中的 Service，可用于处理一些通用的数据库操作。GetPage 方法用于获取分页数据，根据请求参数进行条件筛选和权限控制，将数据分页返回。Get 方法用于获取单条数据，根据请求参数进行权限控制，将数据返回。Insert 方法用于插入数据，将请求参数转换为数据模型，存储到数据库中。Update 方法用于更新数据，先根据请求参数获取要更新的数据，再将请求参数转换为数据模型，更新到数据库中。Remove 方法用于删除数据，先根据请求参数获取要删除的数据，再根据权限控制删除数据。
// 在这些方法中，都使用了 actions 包中的 DataPermission 进行权限控制。同时使用了 dto 中的请求参数转换为数据模型的方法，将请求参数和数据库表中的字段进行对应，方便进行数据库操作。
