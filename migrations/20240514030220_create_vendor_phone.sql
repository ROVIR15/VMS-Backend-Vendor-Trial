-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vendor_phone_type') THEN
        CREATE TYPE vendor_phone_type AS ENUM ('work', 'personal');
    END IF;
END $$;
CREATE TABLE public."VendorPhone" (
    "id" serial PRIMARY KEY,
	"vendor_id" int4 NULL,
	"phone_country_code" varchar(5) NULL,
    "phone_number" varchar(255) NULL,
    "vendor_phone_type" public.vendor_phone_type NULL,
	CONSTRAINT vendorphone_vendors_fk FOREIGN KEY ("vendor_id") REFERENCES public."Vendors"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public."vendor_phone_type";
ALTER TABLE public."VendorPhone" DROP CONSTRAINT IF EXISTS vendorphone_vendors_fk;
DROP TABLE IF EXISTS public."VendorPhone";
-- +goose StatementEnd
