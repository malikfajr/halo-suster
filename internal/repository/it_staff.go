package repository

import (
	"context"
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
)

type ITRepository interface {
	Create(ctx context.Context, pool *pgxpool.Pool, ITStaff *entity.ITStaff) error
	Login(ctx context.Context, pool *pgxpool.Pool, ITStaff *entity.ITStaff) error
	NipIsExist(ctx context.Context, pool *pgxpool.Pool, nip int) bool
}

type itRepository struct{}

func NewItRepository() ITRepository {
	return &itRepository{}
}

func (i *itRepository) Create(ctx context.Context, pool *pgxpool.Pool, ITStaff *entity.ITStaff) error {
	query := "INSERT INTO users(id, nip, name, password, role) VALUES ($1, $2, $3, $4, 1)"

	_, err := pool.Exec(ctx, query, ITStaff.ID, strconv.Itoa(ITStaff.Nip), ITStaff.Name, ITStaff.Password)
	if err != nil {
		return err
	}

	return nil
}

// Seacrh it account by nip
func (i *itRepository) Login(ctx context.Context, pool *pgxpool.Pool, ITStaff *entity.ITStaff) error {
	query := "SELECT id, name, password FROM users WHERE nip = $1 LIMIT 1;"

	err := pool.QueryRow(ctx, query, strconv.Itoa(ITStaff.Nip)).Scan(&ITStaff.ID, &ITStaff.Name, &ITStaff.Password)
	if err != nil {
		return errors.New("nip not found")
	}

	return nil
}

// Search nip if exist
func (i *itRepository) NipIsExist(ctx context.Context, pool *pgxpool.Pool, nip int) bool {
	var exist int
	query := "SELECT 1 FROM users WHERE nip = $1 LIMIT 1;"

	err := pool.QueryRow(ctx, query, strconv.Itoa(nip)).Scan(&exist)
	if err != nil {
		return false
	}

	return true
}
