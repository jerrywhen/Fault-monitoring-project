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

type ArticleA struct {
	service.Service
}

// GetPage 获取ArticleA列表
func (e *ArticleA) GetPage(c *dto.ArticleAGetPageReq, p *actions.DataPermission, list *[]models.ArticleA, count *int64) error {
	var err error
	var data models.ArticleA

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ArticleAService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ArticleA对象
func (e *ArticleA) Get(d *dto.ArticleAGetReq, p *actions.DataPermission, model *models.ArticleA) error {
	var data models.ArticleA

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetArticleA error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ArticleA对象
func (e *ArticleA) Insert(c *dto.ArticleAInsertReq) error {
    var err error
    var data models.ArticleA
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ArticleAService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ArticleA对象
func (e *ArticleA) Update(c *dto.ArticleAUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.ArticleA{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("ArticleAService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除ArticleA
func (e *ArticleA) Remove(d *dto.ArticleADeleteReq, p *actions.DataPermission) error {
	var data models.ArticleA

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveArticleA error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
