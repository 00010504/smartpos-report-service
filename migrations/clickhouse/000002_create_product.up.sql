CREATE TABLE "product" (
    "id" UUID,
    "company_id" UUID
) ENGINE = MergeTree ORDER BY tuple()
