package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"github.com/okamin-chen/service/app/entity"
	"gorm.io/gorm"
)

type (
	CardService interface {
		Insert(data entity.Card) *entity.Card
		InsertBatch(data []entity.Card) *[]entity.Card
		FindOne(id int64) *entity.Card
		FindListPage(page, pageSize int, data entity.Card) (*[]entity.Card, int64)
		FindList(data entity.Card) *[]entity.Card
		Update(data entity.Card) *entity.Card
		Delete(ids []int64)
	}

	cardServiceImpl struct {
		Table string
		Db    *gorm.DB
	}
)

func NewCardServiceImpl(db *gorm.DB) *cardServiceImpl {
	return &cardServiceImpl{
		Table: entity.Card{}.TableName(),
		Db:    db,
	}
}

func (m *cardServiceImpl) Insert(data entity.Card) *entity.Card {
	err := m.Db.Table(m.Table).Create(&data).Error
	biz.ErrIsNil(err, "添加Card失败")
	return &data
}

func (m *cardServiceImpl) InsertBatch(data []entity.Card) *[]entity.Card {
	err := m.Db.Table(m.Table).CreateInBatches(&data, 100).Error
	biz.ErrIsNil(err, "添加Card失败")
	return &data
}

func (m *cardServiceImpl) FindOne(id int64) *entity.Card {
	resData := new(entity.Card)
	db := m.Db.Table(m.Table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询Card失败")
	return resData
}

func (m *cardServiceImpl) FindListPage(page, pageSize int, data entity.Card) (*[]entity.Card, int64) {
	list := make([]entity.Card, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := m.Db.Table(m.Table)
	// 此处填写 where参数判断
	if data.Title != "" {
		db = db.Where("title = ?", data.Title)
	}
	if data.Code != "" {
		db = db.Where("code = ?", data.Code)
	}
	if data.Bank != "" {
		db = db.Where("bank = ?", data.Bank)
	}
	if data.Organize != "" {
		db = db.Where("organize = ?", data.Organize)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.CardNo != "" {
		db = db.Where("card_no = ?", data.CardNo)
	}
	db.Order("id desc")
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询Card分页列表失败")
	return &list, total
}

func (m *cardServiceImpl) FindList(data entity.Card) *[]entity.Card {
	list := make([]entity.Card, 0)
	db := m.Db.Table(m.Table)
	// 此处填写 where参数判断
	if data.Title != "" {
		db = db.Where("title = ?", data.Title)
	}
	if data.Code != "" {
		db = db.Where("code = ?", data.Code)
	}
	if data.Bank != "" {
		db = db.Where("bank = ?", data.Bank)
	}
	if data.Organize != "" {
		db = db.Where("organize = ?", data.Organize)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.CardNo != "" {
		db = db.Where("card_no = ?", data.CardNo)
	}
	db.Order("id desc")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询Card列表失败")
	return &list
}

func (m *cardServiceImpl) Update(data entity.Card) *entity.Card {
	biz.ErrIsNil(m.Db.Table(m.Table).Updates(&data).Error, "修改Card失败")
	return &data
}

func (m *cardServiceImpl) Delete(ids []int64) {
	biz.ErrIsNil(m.Db.Table(m.Table).Delete(&entity.Card{}, "id in (?)", ids).Error, "删除Card失败")
}
