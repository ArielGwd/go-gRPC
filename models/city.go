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

// create
func (c *City) Create(ctx context.Context, db *sql.DB, in *cities.CityInput) error {
	query := `INSERT INTO cities (name) VALUES ($1) RETURNING id`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	err = stmt.QueryRowContext(ctx, in.Name).Scan(&c.Pb.Id)
	if err != nil {
		return err
	}

	c.Pb.Name = in.Name

	return nil
}

// delete
func (c *City) Delete(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `DELETE FROM cities WHERE id = $1`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, in.Id)
	if err != nil {
		return err
	}

	return nil
}

// update
func (c *City) Update(ctx context.Context, db *sql.DB, in *cities.City) error {
	query := `UPDATE cities SET name = $2 WHERE id = $1`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, in.Name, in.Id)
	if err != nil {
		return err
	}

	c.Pb.Id = in.Id
	c.Pb.Name = in.Name

	return nil
}
