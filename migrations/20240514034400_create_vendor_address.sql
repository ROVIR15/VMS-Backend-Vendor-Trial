-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public."VendorAddress" (
	"id" serial PRIMARY KEY,
	"vendor_id" int NULL,
	"address" text NULL,
	"location" GEOMETRY(POINT,4326),
	"is_primary" bool DEFAULT false NULL,
	CONSTRAINT vendoraddress_vendors_fk FOREIGN KEY ("vendor_id") REFERENCES public."Vendors"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public."VendorAddress" DROP CONSTRAINT IF EXISTS vendoraddress_vendors_fk;
DROP TABLE IF EXISTS public."VendorAddress";
-- +goose StatementEnd
