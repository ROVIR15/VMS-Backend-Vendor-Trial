-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public."VendorBookingPolicy" (
	"id" serial PRIMARY KEY,
	"vendor_id" int4 NULL,
	"days_prior" INT4 DEFAULT 0 NOT NULL,
	CONSTRAINT vendorbookingpolicy_unique UNIQUE ("vendor_id"),
	CONSTRAINT vendorbookingpolicy_vendors_fk FOREIGN KEY ("vendor_id") REFERENCES public."Vendors"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX vendorbookingpolicy_vendor_id_idx ON public."VendorBookingPolicy" USING btree ("vendor_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public."VendorBookingPolicy" DROP CONSTRAINT IF EXISTS vendorbookingpolicy_vendors_fk;
DROP TABLE IF EXISTS public."VendorBookingPolicy";
-- +goose StatementEnd
