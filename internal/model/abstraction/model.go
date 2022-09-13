package abstraction

import (
	"time"

	"github.com/mcholismalik/boilerplate-golang/pkg/util/date"

	"gorm.io/gorm"
)

type Entity struct {
	ID string `json:"id" gorm:"primaryKey;"`

	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`

	DeletedAt *time.Time `json:"-" gorm:"index"`
}

func (m *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}

func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}
