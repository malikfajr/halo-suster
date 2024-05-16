package repository

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
)

type NurseRepo interface {
	NipIsExist(ctx context.Context, pool *pgxpool.Pool, nip int) bool
	Create(ctx context.Context, pool *pgxpool.Pool, payload *entity.AddNursePayload) error
}

type nurseRepo struct{}

// TODO: fix validation

// Create implements NurseRepo.
func (*nurseRepo) Create(ctx context.Context, pool *pgxpool.Pool, payload *entity.AddNursePayload) error {
	query := "INSERT INTO users(id, nip, name, card_img, role) VALUES($1, $2, $3, $4, 0);"

	_, err := pool.Exec(ctx, query, payload.Id, strconv.Itoa(payload.Nip), payload.Name, payload.IdentityCardScanImg)
	if err != nil {
		return err
	}

	return nil
}

// Search nip if exist
func (n *nurseRepo) NipIsExist(ctx context.Context, pool *pgxpool.Pool, nip int) bool {
	var exist int
	query := "SELECT 1 FROM users WHERE nip = $1 LIMIT 1;"

	err := pool.QueryRow(ctx, query, strconv.Itoa(nip)).Scan(&exist)
	if err != nil {
		return false
	}

	return true
}

func NewNurseRepo() NurseRepo {
	return &nurseRepo{}
}
