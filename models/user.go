package models

import (
	"context"
	"database/sql"
	"proyek/pb/cities"
)

type City struct {
	Pb cities.City
}

func (c *City) Get(ctx context.Context, db *sql.DB, in *cities.Id) error {
	c.Pb.Id = 1
	c.Pb.Name = "Bandung"

	query := `SELECT id, name FROM cities WHERE id = $1`
	err := db.QueryRowContext(ctx, query, in.Id).Scan(&c.Pb.Id, &c.Pb.Name)

	if err != nil {
		return err
	}

	return nil
}
