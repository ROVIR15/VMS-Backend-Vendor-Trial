-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vendor_visibility') THEN
        CREATE TYPE vendor_visibility AS enum ('anyone', 'invited_host');
    END IF;
END $$;

CREATE TABLE public."Vendors" (
    "id" serial PRIMARY KEY,
    "name" varchar(255) NULL,
    "email" varchar(255) NULL,
    "commission" float8 NULL,
    "vendor_visibility" public."vendor_visibility" NULL,
    "description" varchar(255) NULL,
    "logo" varchar(255) NULL,
    "trial_end_date_at" timestamp NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL,
    "deleted_at" timestamp NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public."vendor_visibility";
DROP TABLE IF EXISTS public."Vendors";
-- +goose StatementEnd
