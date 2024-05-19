CREATE TABLE IF NOT EXISTS medical_record(
    id BIGSERIAL PRIMARY KEY,
    patient_id CHAR(16) NOT NULL,
    user_id CHAR(36) NOT NULL,
    user_nip VARCHAR(15) NOT NULL,
    symptoms TEXT NOT NULL,
    medications TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);


CREATE INDEX idx_medical_record_patient_id ON medical_record(patient_id);
