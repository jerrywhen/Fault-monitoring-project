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

type FyUser struct {
	service.Service
}

// GetPage 获取FyUser列表
func (e *FyUser) GetPage(c *dto.FyUserGetPageReq, p *actions.DataPermission, list *[]models.FyUser, count *int64) error {
	var err error
	var data models.FyUser

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FyUserService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FyUser对象
func (e *FyUser) Get(d *dto.FyUserGetReq, p *actions.DataPermission, model *models.FyUser) error {
	var data models.FyUser

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFyUser error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FyUser对象
func (e *FyUser) Insert(c *dto.FyUserInsertReq) error {
	var err error
	var data models.FyUser
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FyUserService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FyUser对象
func (e *FyUser) Update(c *dto.FyUserUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FyUser{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("FyUserService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FyUser
func (e *FyUser) Remove(d *dto.FyUserDeleteReq, p *actions.DataPermission) error {
	var data models.FyUser

	db := e.Orm.Unscoped().Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFyUser error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

//这段代码定义了一个FyUser服务结构体，它继承了sdk/service包中的Service结构体。其中定义了GetPage()、Get()、Insert()、Update()和Remove()五个方法，分别用于获取FyUser列表、获取FyUser对象、创建FyUser对象、修改FyUser对象和删除FyUser对象。这些方法都接受一个或多个参数，例如GetPage()方法接受FyUserGetPageReq类型的参数c、DataPermission类型的参数p，以及一个FyUser类型的指针list和一个int64类型的指针count。其中，FyUserGetPageReq是一个结构体，它定义了查询参数，例如分页、搜索等；DataPermission是一个结构体，它定义了数据权限相关的操作，例如过滤掉没有权限的数据。这些方法都会调用e.Orm提供的方法，例如Model()、Scopes()、Find()、First()、Create()和Save()等。它们还会根据具体业务逻辑进行一些额外的操作，例如根据ID获取对象、设置创建人和更新人的ID等。如果发生错误，它们会返回一个错误对象。
