package repository

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
)

type ITRepository interface {
	Create(ctx context.Context, pool *pgxpool.Pool, ITStaff *entity.ITStaff) error
	Login(ctx context.Context, pool *pgxpool.Pool, ITStaff *entity.ITStaff) error
	NipIsExist(ctx context.Context, pool *pgxpool.Pool, nip int) bool
	GetAllUsers(ctx context.Context, pool *pgxpool.Pool, params *entity.UserParam) []*entity.User
}

type itRepository struct{}

func NewItRepository() ITRepository {
	return &itRepository{}
}

// TODO: Fix validation

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

func (i *itRepository) GetAllUsers(ctx context.Context, pool *pgxpool.Pool, params *entity.UserParam) []*entity.User {
	var users []*entity.User

	query := "SELECT id, nip::BIGINT, name, created_at FROM users WHERE 1=1"
	args := pgx.NamedArgs{}

	if params.UserId != "" {
		query += " AND id = @user_id"
		args["user_id"] = params.UserId
	}

	if params.Name != "" {
		query += " AND name LIKE @name"
		args["name"] = "%" + strings.ToLower(params.Name) + "%"
	}

	if params.Nip != "" {
		query += " AND nip LIKE @nip "
		args["nip"] = params.Nip + "%"
	}

	if params.Role != "" {
		query += " AND role = @role"
		args["role"] = i.GetRoleIndex(params.Role)
	}

	query += " OFFSET @offset LIMIT @limit"
	args["offset"] = params.Offset
	args["limit"] = params.Limit

	rows, err := pool.Query(ctx, query, args)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		user := &entity.User{}
		rows.Scan(&user.ID, &user.Nip, &user.Name, &user.CreatedAt)
		users = append(users, user)
		log.Println(user)
	}

	return users
}

func (i *itRepository) GetRoleIndex(role string) int {
	roles := map[string]int{
		"it":    1,
		"nurse": 0,
	}

	j, _ := roles[role]
	return j
}
