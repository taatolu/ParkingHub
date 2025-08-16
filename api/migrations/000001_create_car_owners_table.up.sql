-- 車両所有者テーブルの作成
CREATE TABLE car_owners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    license_number VARCHAR(50) UNIQUE NOT NULL,
    license_expiry DATE NOT NULL,
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- インデックス作成
CREATE INDEX idx_car_owners_email ON car_owners(email);
CREATE INDEX idx_car_owners_license ON car_owners(license_number);

-- コメント追加
COMMENT ON TABLE car_owners IS '車両所有者情報';
COMMENT ON COLUMN car_owners.license_expiry IS '運転免許証有効期限';
