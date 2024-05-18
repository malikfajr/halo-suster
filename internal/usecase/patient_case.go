package usecase

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
	"github.com/malikfajr/halo-suster/internal/exception"
	"github.com/malikfajr/halo-suster/internal/repository"
)

type PatientCase struct {
	pool        *pgxpool.Pool
	patientRepo *repository.PatientRepository
}

func NewPatientCase(pool *pgxpool.Pool, patientRepo *repository.PatientRepository) *PatientCase {
	return &PatientCase{
		pool:        pool,
		patientRepo: patientRepo,
	}
}

func (p *PatientCase) Insert(ctx context.Context, payload *entity.AddPatientPayload) (*entity.Patient, error) {
	patient := &entity.Patient{
		IdNumber:    payload.IdNumber,
		PhoneNumber: payload.PhoneNumber,
		Name:        payload.Name,
		BirthDate:   payload.BirthDate,
		Gender:      payload.Gender,
		ImageCard:   payload.ImageCard,
	}

	if exist := p.patientRepo.IdNumberExis(ctx, p.pool, payload.IdNumber); exist == true {
		return nil, exception.NewConflict("identity number is exist")
	}

	p.patientRepo.Insert(ctx, p.pool, patient)

	return patient, nil
}

func (p *PatientCase) GetAll(ctx context.Context, params *entity.PatientQueryParam) []entity.Patient {

	if p.isValidOrder(params.CreatedAt) == false {
		params.CreatedAt = ""
	}

	if params.Limit <= 0 {
		params.Limit = 5
	}

	if params.Offset < 0 {
		params.Offset = 0
	}

	patients := p.patientRepo.GetAll(ctx, p.pool, params)
	return patients
}

func (p *PatientCase) isValidOrder(key string) bool {
	orders := map[string]bool{
		"asc":  true,
		"desc": true,
	}

	_, ok := orders[key]
	return ok
}
