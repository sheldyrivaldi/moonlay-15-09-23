package repository

import (
	"fmt"
	"moonlay-todolist/internal/abstraction"
	"moonlay-todolist/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ISubListRepository interface {
	FindAll(d *string, t *string, p *abstraction.Pagination) (*[]model.SubList, *abstraction.PaginationInfo, error)
	FindByID(ID string) (*model.SubList, error)
	Create(data *model.SubList) (*model.SubList, error)
	CreateFile(link string, subListID string) (*model.SubListFile, error)
	UpdateByID(ID string, data *model.SubList) (*model.SubList, error)
	DeleteByID(ID string) (*model.SubList, error)
	DeleteFileBySubListID(subListID string) error
}

type SubListRepository struct {
	db *gorm.DB
}

func NewSubListRepository(db *gorm.DB) *SubListRepository {
	return &SubListRepository{db}
}

func (r *SubListRepository) FindAll(d *string, t *string, p *abstraction.Pagination) (*[]model.SubList, *abstraction.PaginationInfo, error) {
	var datas []model.SubList
	var info abstraction.PaginationInfo

	query := r.db.Model(&model.SubList{})

	// sort
	if p.Sort == nil {
		sort := "desc"
		p.Sort = &sort
	}
	if p.SortBy == nil {
		sortBy := "created_at"
		p.SortBy = &sortBy
	}

	sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

	// pagination
	if p.Page == nil {
		page := 1
		p.Page = &page
	}
	if p.PageSize == nil {
		pageSize := 10
		p.PageSize = &pageSize
	}

	info = abstraction.PaginationInfo{
		Pagination: p,
	}

	limit := *p.PageSize + 1
	offset := (*p.Page - 1) * limit

	query = query.Limit(limit).Offset(offset)

	if t != nil && d != nil {
		query = query.Where("title LIKE ?", "%"+*t+"%").Where("description LIKE ?", "%"+*d+"%")
	} else if t != nil {
		query = query.Where("title LIKE ?", "%"+*t+"%")
	} else if d != nil {
		query = query.Where("description LIKE ?", "%"+*d+"%")
	}

	err := query.Preload("Files").Find(&datas).Error

	if err != nil {
		return &datas, &info, err
	}

	info.Count = len(datas)
	info.MoreRecords = false
	if len(datas) > *p.PageSize {
		info.MoreRecords = true
		info.Count -= 1
		datas = datas[:len(datas)-1]
	}

	return &datas, &info, nil
}

func (r *SubListRepository) FindByID(ID string) (*model.SubList, error) {
	var data *model.SubList

	err := r.db.Preload("Files").Where("id = ?", ID).First(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *SubListRepository) Create(data *model.SubList) (*model.SubList, error) {

	err := r.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(data).Preload("Files").First(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *SubListRepository) UpdateByID(ID string, data *model.SubList) (*model.SubList, error) {
	var list model.SubList
	err := r.db.Preload("Files").Where("id = ?", ID).First(&list).Error
	if err != nil {
		return nil, err
	}

	if data.Title != "" {
		list.Title = data.Title
	}
	if data.Description != "" {
		list.Description = data.Description
	}

	err = r.db.Save(&list).Error
	if err != nil {
		return nil, err
	}

	return &list, nil
}

func (r *SubListRepository) DeleteByID(ID string) (*model.SubList, error) {
	var data *model.SubList
	err := r.db.Preload("Files").Where("id = ?", ID).First(&data).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Delete(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r SubListRepository) CreateFile(link string, subListID string) (*model.SubListFile, error) {
	data := &model.SubListFile{}
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	data.ID = ID.String()
	data.SubListID = subListID
	data.Link = link

	err = r.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(data).First(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *SubListRepository) DeleteFileBySubListID(subListID string) error {
	err := r.db.Where("sub_list_id = ?", subListID).Delete(&model.SubListFile{}).Error
	return err
}
