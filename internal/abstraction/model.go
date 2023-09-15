package abstraction

import "time"

type Entity struct {
	ID string `json:"id" gorm:"primaryKey"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
