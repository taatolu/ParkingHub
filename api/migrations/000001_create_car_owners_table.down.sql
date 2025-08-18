-- インデックス削除
DROP INDEX IF EXISTS idx_car_owners_email;
DROP INDEX IF EXISTS idx_car_owners_license;

-- 制約の削除
ALTER TABLE car_owners DROP CONSTRAINT IF EXISTS car_owners_pkey;
ALTER TABLE car_owners DROP CONSTRAINT IF EXISTS car_owners_email_key;
ALTER TABLE car_owners DROP CONSTRAINT IF EXISTS car_owners_license_number_key;

-- 他に必要なら追加
DROP TABLE IF EXISTS car_owners;