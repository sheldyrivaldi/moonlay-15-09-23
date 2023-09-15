package repository

import (
	"fmt"
	"moonlay-todolist/internal/abstraction"
	"moonlay-todolist/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IListRepository interface {
	FindAll(s *bool, d *string, t *string, p *abstraction.Pagination) (*[]model.List, *abstraction.PaginationInfo, error)
	FindByID(ID string) (*model.List, error)
	Create(data *model.List) (*model.List, error)
	CreateFile(link string, listID string) (*model.ListFile, error)
	UpdateByID(ID string, data *model.List) (*model.List, error)
	DeleteByID(ID string) (*model.List, error)
	DeleteFileByListID(ListID string) error
}

type ListRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) *ListRepository {
	return &ListRepository{db}
}

func (r *ListRepository) FindAll(s *bool, d *string, t *string, p *abstraction.Pagination) (*[]model.List, *abstraction.PaginationInfo, error) {
	var datas []model.List
	var info abstraction.PaginationInfo

	query := r.db.Model(&model.List{})

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

	if *s {
		query = query.Preload("Files").Preload("SubList").Preload("SubList.Files")
	}

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

func (r *ListRepository) FindByID(ID string) (*model.List, error) {
	var data model.List

	err := r.db.Preload("Files").Preload("SubList").Preload("SubList.Files").Where("id = ?", ID).First(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ListRepository) Create(data *model.List) (*model.List, error) {

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

func (r *ListRepository) UpdateByID(ID string, data *model.List) (*model.List, error) {
	var list model.List
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

func (r *ListRepository) DeleteByID(ID string) (*model.List, error) {
	var data *model.List
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

func (r ListRepository) CreateFile(link string, listID string) (*model.ListFile, error) {
	data := &model.ListFile{}
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	data.ID = ID.String()
	data.ListID = listID
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

func (r *ListRepository) DeleteFileByListID(ListID string) error {
	err := r.db.Where("list_id = ?", ListID).Delete(&model.ListFile{}).Error
	return err
}
