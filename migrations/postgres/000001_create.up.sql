CREATE TYPE user_type AS ENUM (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', -- system user
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'  -- admin user
);

CREATE TYPE app_type AS ENUM (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', -- client
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'  -- admin
);

CREATE TABLE IF NOT EXISTS "user" (
    "id" UUID PRIMARY KEY,
    "user_type_id" user_type NOT NULL,
    "first_name" VARCHAR(250) NOT NULL,
    "last_name" VARCHAR(250) NOT NULL,
    "phone_number" VARCHAR(30) NOT NULL,
    "image" TEXT,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX "user_deleted_at_idx" ON "user"("deleted_at");

INSERT INTO "user" (
    "id",
    "first_name",
    "last_name",
    "phone_number",
    "user_type_id"
) VALUES (
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    'admin',
    'admin',
    '99894172774',
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'
);


CREATE TABLE IF NOT EXISTS "company" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "created_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX company_deleted_at_idx ON "company"("deleted_at");

CREATE TABLE IF NOT EXISTS "company_user" (
    "user_id" UUID NOT NULL REFERENCES "user"("id"),
    "company_id" UUID NOT NULL REFERENCES "company"("id"),
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY("user_id", "company_id", "deleted_at")
);

CREATE INDEX "company_user_deleted_at_idx" ON "company_user"("deleted_at");


CREATE TABLE IF NOT EXISTS "shop" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "company_id" UUID NOT NULL,
    "created_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX shop_deleted_at_idx ON "shop"("deleted_at");

CREATE TABLE "product" (
  "id" UUID PRIMARY KEY,
  "is_marking" BOOLEAN NOT NULL DEFAULT FALSE,
  "sku" VARCHAR NOT NULL,
  "measurement_unit" UUID,
  "mxik_code" VARCHAR,
  "name" varchar NOT NULL,
  "company_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("name", "company_id", "deleted_at")
);
CREATE INDEX "product_deleted_at_idx" ON "product"("deleted_at");

