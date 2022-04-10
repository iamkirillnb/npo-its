package repos

import (
	"npo-its/internal/entities"
	"npo-its/pkg/pg"
)


type DBRepo struct {
	*pg.DB
}

func NewDB(d *pg.DB) *DBRepo {
	return &DBRepo{d}
}

func (d *DBRepo) GetFiveMaxMetrics() ([]*entities.Item, error) {
	const qry = `SELECT * FROM item6 ORDER BY metric DESC LIMIT 5`

	var data []*entities.Item

	err := d.DB.Select(&data, qry)
	if err != nil {
		return nil, err
	}

	return data, nil
}
