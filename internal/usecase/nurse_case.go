package usecase

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
	"github.com/malikfajr/halo-suster/internal/exception"
	"github.com/malikfajr/halo-suster/internal/repository"
)

type NurseCase interface {
	// 	Login(e echo.Context) error
	Register(ctx context.Context, payload *entity.AddNursePayload) error
	// GetById(e echo.Context) error
	// Destroy(e echo.Context) error
	// AddAccess(e echo.Context) error
}

type nurseCase struct {
	pool      *pgxpool.Pool
	nurseRepo repository.ITRepository
}

// Register implements NurseCase.
func (n *nurseCase) Register(ctx context.Context, payload *entity.AddNursePayload) error {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	payload.Id = id.String()

	nRepo := repository.NewNurseRepo()

	if exist := nRepo.NipIsExist(ctx, n.pool, payload.Nip); exist == true {
		return exception.NewConflict("nip is exist")
	}

	err = nRepo.Create(ctx, n.pool, payload)
	if err != nil {
		log.Println(err)
		return exception.NewInternalError("internal server error")
	}
	return nil
}

func NewNurseCase(pool *pgxpool.Pool) NurseCase {
	return &nurseCase{
		pool: pool,
	}
}
