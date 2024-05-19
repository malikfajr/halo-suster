package repository

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malikfajr/halo-suster/entity"
)

type medicalRecordRepo struct {
}

type MedicalRecordRepo interface {
	Insert(ctx context.Context, pool *pgxpool.Pool, payload *entity.AddMedicalRecordPayload) error
	GetAll(ctx context.Context, pool *pgxpool.Pool, params *entity.MedicalRecordQueryParam) []entity.MedicalRecord
}

func NewMedicalRecordRepo() MedicalRecordRepo {
	return &medicalRecordRepo{}
}

func (m *medicalRecordRepo) Insert(ctx context.Context, pool *pgxpool.Pool, payload *entity.AddMedicalRecordPayload) error {
	query := "INSERT INTO medical_record(patient_id, user_id, user_nip, symptoms, medications) VALUES(@patientId, @userId, @userNip, @symptoms, @medications)"

	args := pgx.NamedArgs{
		"patientId":   strconv.Itoa(payload.PatientId),
		"userId":      payload.UserId,
		"userNip":     strconv.Itoa(payload.UserNip),
		"symptoms":    payload.Symptoms,
		"medications": payload.Medications,
	}

	_, err := pool.Exec(ctx, query, args)
	if err != nil {
		panic(err)
	}

	return nil
}

func (m *medicalRecordRepo) GetAll(ctx context.Context, pool *pgxpool.Pool, params *entity.MedicalRecordQueryParam) []entity.MedicalRecord {
	query := `SELECT mc.symptoms, mc.medications, mc.created_at,
				json_build_object(
					'nip', u.nip::BIGINT,
					'userId', u.id,
					'name', u.name
				) AS created_by,
				json_build_object(
					'identityNumber', p.id::BIGINT,
					'phoneNumber', p.phone_number,
					'name', p.name,
					'birthDate', p.birth_date,
					'gender', p.gender,
					'identityCardScanImg', p.card_img
				) AS identityDetail
				FROM medical_record mc
				JOIN users u ON mc.user_nip = u.nip
				JOIN patients p ON mc.patient_id = p.id
				AND 1=1`

	args := pgx.NamedArgs{}

	if params.UserId != "" {
		query += " AND u.id = @userId"
		args["userId"] = params.UserId
	}

	if params.PatientId != "" {
		query += " AND p.id = @patientId"
		args["patientId"] = params.PatientId
	}

	if params.UserNip != "" {
		query += " AND mc.nip = @nip"
		args["nip"] = params.UserNip
	}

	if params.CreatedAt != "" {
		query += " ORDER BY mc.created_at " + params.CreatedAt
	}

	query += " LIMIT @limit OFFSET @offset"
	args["limit"] = params.Limit
	args["offset"] = params.Offset

	rows, err := pool.Query(ctx, query, args)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	records := []entity.MedicalRecord{}
	for rows.Next() {
		record := &entity.MedicalRecord{}
		rows.Scan(&record.Symptoms, &record.Medications, &record.CreatedAt, &record.CreatedBy, &record.IdentityDetail)
		records = append(records, *record)
	}

	return records
}
