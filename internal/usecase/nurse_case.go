package usecase

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
	"github.com/malikfajr/halo-suster/internal/exception"
	"github.com/malikfajr/halo-suster/internal/helper/bcrypt"
	"github.com/malikfajr/halo-suster/internal/repository"
)

type NurseCase interface {
	Login(ctx context.Context, payload *entity.AuthLogin) (*entity.User, error)
	Register(ctx context.Context, payload *entity.AddNursePayload) error
	Update(ctx context.Context, payload *entity.EditNursePayload) error
	// GetById(e echo.Context) error
	Delete(ctx context.Context, userId string) error
	AddAccess(ctx context.Context, userId string, password string) error
}

type nurseCase struct {
	pool *pgxpool.Pool
}

// Login implements NurseCase.
func (n *nurseCase) Login(ctx context.Context, payload *entity.AuthLogin) (*entity.User, error) {
	nRepo := repository.NewNurseRepo()

	user, err := nRepo.GetByNip(ctx, n.pool, payload.Nip)
	if err != nil {
		return nil, exception.NewBadRequest("user not found")
	}

	if user.Role != 0 {
		return nil, exception.NewBadRequest("user not nurse")
	}

	if valid := bcrypt.PasswordIsValid(payload.Password, user.Password); valid == false {
		return nil, exception.NewBadRequest("password wrong")
	}

	return user, nil
}

// AddAccess implements NurseCase.
func (n *nurseCase) AddAccess(ctx context.Context, userId string, password string) error {
	nRepo := repository.NewNurseRepo()

	user, err := nRepo.GetById(ctx, n.pool, userId)
	if err != nil {
		return exception.NewBadRequest("user id not found")
	}

	if user.Role != 0 {
		return exception.NewBadRequest("user id is not nurse")
	}

	hashPassword := bcrypt.CreateHash(password)

	nRepo.SetPassword(ctx, n.pool, userId, hashPassword)

	return nil
}

// Delete implements NurseCase.
func (n *nurseCase) Delete(ctx context.Context, userId string) error {
	nRepo := repository.NewNurseRepo()

	user, err := nRepo.GetById(ctx, n.pool, userId)
	if err != nil {
		return exception.NewNotFound("user id not found")
	}

	if user.Role != 0 {
		return exception.NewBadRequest("user is not nurse")
	}

	err = nRepo.Delete(ctx, n.pool, userId)
	if err != nil {
		panic(err)
	}

	return nil
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

// Update implements NurseCase.
func (n *nurseCase) Update(ctx context.Context, payload *entity.EditNursePayload) error {

	nRepo := repository.NewNurseRepo()

	if exist := nRepo.NipIsExist(ctx, n.pool, payload.Nip); exist == true {
		return exception.NewConflict("nip is exist")
	}

	user, err := nRepo.GetById(ctx, n.pool, payload.UserId)
	if err != nil {
		return exception.NewNotFound("User Id not found")
	}

	if user.Role != 0 {
		return exception.NewNotFound("User is not nurse")
	}

	if exist := nRepo.NipIsExist(ctx, n.pool, payload.Nip); exist == true {
		return exception.NewConflict("nip is exists")
	}

	nRepo.Update(ctx, n.pool, payload)

	return nil
}

func NewNurseCase(pool *pgxpool.Pool) NurseCase {
	return &nurseCase{
		pool: pool,
	}
}
