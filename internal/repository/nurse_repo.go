package repository

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
)

type NurseRepo interface {
	NipIsExist(ctx context.Context, pool *pgxpool.Pool, nip int) bool
	Create(ctx context.Context, pool *pgxpool.Pool, payload *entity.AddNursePayload) error
	GetById(ctx context.Context, pool *pgxpool.Pool, userId string) (*entity.User, error)
	GetByNip(ctx context.Context, pool *pgxpool.Pool, nip int) (*entity.User, error)
	Update(ctx context.Context, pool *pgxpool.Pool, payload *entity.EditNursePayload)
	Delete(ctx context.Context, pool *pgxpool.Pool, userId string) error
	SetPassword(ctx context.Context, pool *pgxpool.Pool, userId string, hashPassword string)
}

type nurseRepo struct{}

// GetByNip implements NurseRepo.
func (n *nurseRepo) GetByNip(ctx context.Context, pool *pgxpool.Pool, nip int) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, name, nip::BIGINT, role, password FROM users WHERE nip = $1"

	err := pool.QueryRow(ctx, query, strconv.Itoa(nip)).Scan(&user.ID, &user.Name, &user.Nip, &user.Role, &user.Password)
	if err != nil {
		log.Println(err)
		return nil, errors.New("nip not found")
	}
	return user, nil
}

// SetPassword implements NurseRepo.
func (*nurseRepo) SetPassword(ctx context.Context, pool *pgxpool.Pool, userId string, hashPassword string) {
	query := "UPDATE users SET password = $1 WHERE id = $2"

	_, err := pool.Exec(ctx, query, hashPassword, userId)
	if err != nil {
		panic(err)
	}
}

// Delete implements NurseRepo.
func (*nurseRepo) Delete(ctx context.Context, pool *pgxpool.Pool, userId string) error {
	query := "DELETE FROM users WHERE id = $1"

	_, err := pool.Exec(ctx, query, userId)
	if err != nil {
		panic(err)
	}

	return nil
}

// Update implements NurseRepo.
func (*nurseRepo) Update(ctx context.Context, pool *pgxpool.Pool, payload *entity.EditNursePayload) {

	query := "UPDATE users SET nip = $1, name = $2 WHERE id = $3"

	_, err := pool.Exec(ctx, query, strconv.Itoa(payload.Nip), payload.Name, payload.UserId)
	if err != nil {
		panic(err)
	}

}

// GetById implements NurseRepo.
func (*nurseRepo) GetById(ctx context.Context, pool *pgxpool.Pool, userId string) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, nip::BIGINT, name, role FROM users WHERE id = $1 LIMIT 1"

	err := pool.QueryRow(ctx, query, userId).Scan(&user.ID, &user.Nip, &user.Name, &user.Role)
	if err != nil {
		log.Println(err)
		return nil, errors.New("userId not found")
	}

	return user, nil
}

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
