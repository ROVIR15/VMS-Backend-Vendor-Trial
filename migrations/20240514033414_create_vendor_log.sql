-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vendor_log_section') THEN
        CREATE TYPE vendor_log_section AS ENUM ('vendor', 'vendor_address', 'vendor_phone');
    END IF;
END $$;
CREATE TABLE IF NOT EXISTS public."VendorLog" (
	"id" serial PRIMARY KEY,
	"vendor_id" int4 NOT NULL,
	"user_id" int4 NULL,
	"vendor_log_section" public."vendor_log_section" NULL,
	"old_value" jsonb NULL,
	"new_value" jsonb NULL,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	"deleted_at" timestamp NULL,
	CONSTRAINT vendorlogVendorsFk FOREIGN KEY ("vendor_id") REFERENCES public."Vendors"("id")
);
CREATE INDEX vendorlog_vendor_id_idx ON public."VendorLog" USING btree ("vendor_id", "user_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public."vendor_log_section";
ALTER TABLE public."VendorLog" DROP CONSTRAINT IF EXISTS vendorlog_vendors_fk;
DROP TABLE IF EXISTS public."VendorLog";
-- +goose StatementEnd
