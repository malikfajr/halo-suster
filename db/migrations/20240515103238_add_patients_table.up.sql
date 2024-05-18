CREATE TABLE IF NOT EXISTS patients (
    id CHAR(16) PRIMARY KEY,
    phone_number VARCHAR(15) NOT NULL,
    name VARCHAR(30) NOT NULL,
    birth_date VARCHAR(10) NOT NULL,
    gender VARCHAR(6) NOT NULL,
    card_img TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_patients_male ON patients(gender) WHERE gender = 'male';
CREATE INDEX IF NOT EXISTS idx_patients_female ON patients(gender) WHERE gender = 'female';
CREATE INDEX IF NOT EXISTS idx_patients_name ON patients(LOWER(name));