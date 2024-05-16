package usecase

import (
	"context"
	"log"
	"math"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
	"github.com/malikfajr/halo-suster/internal/exception"
	"github.com/malikfajr/halo-suster/internal/helper/bcrypt"
	"github.com/malikfajr/halo-suster/internal/repository"
)

type ITUseCase interface {
	Register(ctx context.Context, payload *entity.ITStaffRegister) (*entity.ITStaff, error)
	Login(ctx context.Context, payload *entity.ITStaffLogin) (*entity.ITStaff, error)
}

type iTUsecase struct {
	pool         *pgxpool.Pool
	itRepository repository.ITRepository
}

// Login implements ITUseCase.
func (i *iTUsecase) Login(ctx context.Context, payload *entity.ITStaffLogin) (*entity.ITStaff, error) {
	user := &entity.ITStaff{
		Nip: payload.Nip,
	}

	validPrefix := int(payload.Nip / int(math.Pow(10, 10))) // check prefix for it nip

	err := i.itRepository.Login(ctx, i.pool, user)
	if err != nil || validPrefix != 615 {
		log.Println(validPrefix)
		return nil, exception.NewNotFound("nip not found")
	}

	return user, nil
}

// Register implements iITUseCase.
func (i *iTUsecase) Register(ctx context.Context, payload *entity.ITStaffRegister) (*entity.ITStaff, error) {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	it_staff := &entity.ITStaff{
		ID:       id.String(),
		Nip:      payload.Nip,
		Name:     payload.Name,
		Password: bcrypt.CreateHash(payload.Password),
	}

	if exist := i.itRepository.NipIsExist(ctx, i.pool, payload.Nip); exist == true {
		return nil, exception.NewConflict("nip is exist")
	}

	if err := i.itRepository.Create(ctx, i.pool, it_staff); err != nil {
		log.Println(err)
		return nil, exception.NewInternalError("internal server error")
	}

	return it_staff, nil
}

func NewItUsecase(pool *pgxpool.Pool, itRepository repository.ITRepository) ITUseCase {
	return &iTUsecase{pool: pool, itRepository: itRepository}
}
