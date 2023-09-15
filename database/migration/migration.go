package migration

import (
	"fmt"
	"moonlay-todolist/database"
	"moonlay-todolist/internal/model"

	"gorm.io/gorm"
)

type Migration interface {
	AutoMigrate()
	SetDb(*gorm.DB)
}

type migration struct {
	Db            *gorm.DB
	DbModels      *[]interface{}
	IsAutoMigrate bool
}

func Init() {
	mgConfigurations := map[string]Migration{
		"POSTGRES": &migration{
			DbModels: &[]interface{}{
				&model.List{},
				&model.ListFile{},
				&model.SubList{},
				&model.SubListFile{},
			},
			IsAutoMigrate: true,
		},
	}

	for k, v := range mgConfigurations {
		dbConnection, err := database.Connection(k)
		if err != nil {
			fmt.Printf("Failed to run migration, database not found %s", k)
		} else {
			v.SetDb(dbConnection)
			v.AutoMigrate()
			fmt.Printf("Successfully run migration for database %s", k)
		}
	}

}

func (m *migration) AutoMigrate() {
	if m.IsAutoMigrate {
		m.Db.AutoMigrate(*m.DbModels...)
	}
}

func (m *migration) SetDb(db *gorm.DB) {
	m.Db = db
}
