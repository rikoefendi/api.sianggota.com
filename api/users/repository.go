package users

import (
	"api.sianggota.com/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

func New() *Repository {
	db := database.Session()
	r := &Repository{db}
	return r
}

func (r *Repository) Create(u UserCreateInput) (Model, error) {
	m := Model{
		Name: &u.Name, Email: &u.Email, Password: u.Password,
	}
	result := r.db.Create(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

func (r *Repository) UpdateById(id string, dest Model) (Model, error) {
	m := Model{}
	result := r.db.Where("id = ?", id).Clauses(clause.Locking{Strength: "UPDATE"}).First(&m).Updates(dest)
	if result.Error != nil {
		return m, result.Error
	}
	if result.RowsAffected < 1 {
		return m, gorm.ErrRecordNotFound
	}
	return m, nil
}

func (r *Repository) FetchById(id string) (user Model, err error) {
	result := r.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (r *Repository) ShowAll(page int, perPage int) (users []Model, paginated database.Paginated, err error) {
	paginated = database.Paginate(r.db.Where("name like '%s%'")).SetPage(page).SetPerPage(perPage).Exec(&users)
	if paginated.Error != nil {
		return users, paginated, paginated.Error
	}
	return users, paginated, nil
}
