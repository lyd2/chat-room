package models

type Room struct {
	ID          int    `gorm:"primary_key" json:"id" validate:"min=1"`
	Name        string `json:"name" validate:"min=3,max=50"`
	Description string `json:"description" validate:"max=200"`
	Status      string `json:"status" validate:"oneof=0 1"`
	CreatedAt   int    `json:"createdAt"`
	UpdatedAt   int    `json:"updatedAt"`
}

func (r *Room) All() []Room {
	var res []Room
	db.Find(&res)
	return res
}

func (r *Room) Rooms(pageNum, pageSize int, where interface{}) []Room {

	var res []Room
	db.Select("id, name, description, status").
		Where(where).
		Offset(pageNum).
		Limit(pageSize).
		Find(&res)

	return res
}

func (r *Room) Total(where interface{}) int {

	var res int
	db.Model(r).Where(where).Count(&res)

	return res

}

func (r *Room) Info() bool {
	where := Room{
		ID: r.ID,
	}
	r.ID = 0
	db.Where(where).First(r)

	if r.ID > 0 {
		return true
	}
	return false
}
