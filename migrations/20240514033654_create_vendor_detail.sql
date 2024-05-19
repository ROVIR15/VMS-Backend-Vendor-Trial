-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public."VendorDetail" (
	"id" serial PRIMARY KEY,
	"vendor_id" int4 NOT NULL,
	"commission" float4 DEFAULT 0 NOT NULL,
	"is_commission_include_tax" bool DEFAULT true NOT NULL,
	CONSTRAINT vendordetail_unique UNIQUE ("vendor_id"),
	CONSTRAINT vendordetail_vendors_fk FOREIGN KEY ("vendor_id") REFERENCES public."Vendors"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX vendordetail_vendor_id_idx ON public."VendorDetail" USING btree ("vendor_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public."VendorDetail" DROP CONSTRAINT IF EXISTS vendordetail_vendors_fk;
DROP TABLE IF EXISTS public."VendorDetail";
-- +goose StatementEnd
