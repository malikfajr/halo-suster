package usecase

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
	"github.com/malikfajr/halo-suster/internal/exception"
	"github.com/malikfajr/halo-suster/internal/repository"
)

type medicalCase struct {
	pool *pgxpool.Pool
}

func NewMedicalCase(pool *pgxpool.Pool) *medicalCase {
	return &medicalCase{pool: pool}
}

func (m *medicalCase) Insert(ctx context.Context, payload *entity.AddMedicalRecordPayload) error {
	pCase := &repository.PatientRepository{}

	exist := pCase.IdNumberExis(ctx, m.pool, payload.PatientId)
	if exist == false {
		return exception.NewNotFound("identityNumber not exist")
	}

	mRepo := repository.NewMedicalRecordRepo()

	err := mRepo.Insert(ctx, m.pool, payload)
	if err != nil {
		panic(err)
	}

	return nil
}

func (m *medicalCase) GetAll(ctx context.Context, params *entity.MedicalRecordQueryParam) []entity.MedicalRecord {
	if m.isValidOrder(params.CreatedAt) == false {
		params.CreatedAt = ""
	}

	if params.Limit <= 0 {
		params.Limit = 5
	}

	if params.Offset < 0 {
		params.Offset = 0
	}

	mRepo := repository.NewMedicalRecordRepo()
	records := mRepo.GetAll(ctx, m.pool, params)

	return records
}

func (m *medicalCase) isValidOrder(key string) bool {
	orders := map[string]bool{
		"asc":  true,
		"desc": true,
	}

	_, ok := orders[key]
	return ok
}
