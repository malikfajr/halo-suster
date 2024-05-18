CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY NOT NULL,
    nip VARCHAR(15) NOT NULL,
    name VARCHAR(50) NOT NULL,
    password CHAR(60) NULL,
    role SMALLINT NOT NULL,
    card_img TEXT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_users_nip ON users(nip);

CREATE INDEX IF NOT EXISTS idx_users_name ON users(LOWER(name))