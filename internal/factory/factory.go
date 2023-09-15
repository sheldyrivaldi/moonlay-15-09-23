package factory

import (
	"moonlay-todolist/database"
	"moonlay-todolist/internal/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db                *gorm.DB
	ListRepository    repository.IListRepository
	SubListRepository repository.ISubListRepository
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("POSTGRES")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.ListRepository = repository.NewListRepository(f.Db)
	f.SubListRepository = repository.NewSubListRepository(f.Db)
}
