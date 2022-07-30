package database

import (
	"math"

	"gorm.io/gorm"
)

type Paginated struct {
	PerPage   int     `json:"per_page,omitempty"`
	Page      int     `json:"page"`
	Total     int64   `json:"total"`
	PageCount float64 `json:"page_count,omitempty"`
	Error     error   `json:"-"`
	db        *gorm.DB
}

func Paginate(db *gorm.DB) Paginated {
	return Paginated{
		PerPage: 15,
		Page:    1,
		db:      db,
	}
}

func (pg Paginated) SetPage(page int) Paginated {
	if page < 1 {
		return pg
	}
	pg.Page = page
	return pg
}
func (pg Paginated) SetPerPage(perPage int) Paginated {
	if perPage < 1 {
		return pg
	}
	pg.PerPage = perPage
	return pg
}

func (pg Paginated) Exec(dest interface{}) Paginated {
	q := pg.db
	tx := q.Model(dest).Count(&pg.Total)
	if tx.Error != nil {
		pg.Error = tx.Error
		return pg
	}
	pg.PageCount = math.Ceil(float64(pg.Total / int64(pg.PerPage)))
	tx = tx.Order("id asc").Limit(pg.PerPage).Offset((pg.Page - 1) * pg.PerPage).Find(dest)
	if tx.Error != nil {
		pg.Error = tx.Error
	}
	return pg
}
