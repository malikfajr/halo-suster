package repository

import (
	"context"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
)

type PatientRepository struct{}

func (p *PatientRepository) IdNumberExis(ctx context.Context, pool *pgxpool.Pool, idNumber int) bool {
	var exist int
	query := "SELECT 1 FROM patients WHERE id = $1 LIMIT 1;"

	err := pool.QueryRow(ctx, query, strconv.Itoa(idNumber)).Scan(&exist)
	if err != nil {
		return false
	}

	return true
}

func (p *PatientRepository) Insert(ctx context.Context, pool *pgxpool.Pool, patient *entity.Patient) error {
	query := `INSERT INTO 
				patients(id, phone_number, name, birth_date, gender, card_img) 
				VALUES(@id, @phone, @name, @birth, @gender, @img)`

	args := pgx.NamedArgs{
		"id":     strconv.Itoa(patient.IdNumber),
		"name":   patient.Name,
		"phone":  patient.PhoneNumber,
		"birth":  patient.BirthDate,
		"gender": patient.Gender,
		"img":    patient.ImageCard,
	}

	_, err := pool.Exec(ctx, query, args)
	if err != nil {
		panic(err)
	}

	return nil
}

func (p *PatientRepository) GetAll(ctx context.Context, pool *pgxpool.Pool, params *entity.PatientQueryParam) []entity.Patient {

	query := "SELECT id::BIGINT, phone_number, name, birth_date, gender, card_img, created_at FROM patients WHERE 1=1"
	args := pgx.NamedArgs{}

	if params.IdNumber != "" {
		query += " AND id = @id"
		args["id"] = params.IdNumber
	}

	if params.Name != "" {
		query += " AND name LIKE @name"
		args["name"] = "%" + strings.ToLower(params.Name) + "%"
	}

	if params.PhoneNumber != "" {
		query += " AND phone_number LIKE @phone"
		args["phone"] = "%" + params.PhoneNumber + "%"
	}

	if params.CreatedAt != "" {
		query += " ORDER BY created_at " + params.CreatedAt
	}

	query += " LIMIT @limit OFFSET @offset"
	args["limit"] = params.Limit
	args["offset"] = params.Offset

	rows, err := pool.Query(ctx, query, args)
	if err != nil {
		panic(err)
	}

	patients, err := pgx.CollectRows(rows, pgx.RowToStructByPos[entity.Patient])
	if err != nil {
		panic(err)
	}
	return patients
}
