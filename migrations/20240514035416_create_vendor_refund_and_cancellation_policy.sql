-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public."VendorRefundAndCancellationPolicy" (
    "id" serial PRIMARY KEY,
	"vendor_id" int4 NOT NULL,
	"is_guest_pay_transaction" bool DEFAULT true NOT NULL,
	CONSTRAINT vendorrefundandcancellationpolicy_unique UNIQUE ("vendor_id"),
	CONSTRAINT vendorrefundandcancellationpolicy_vendors_fk FOREIGN KEY ("vendor_id") REFERENCES public."Vendors"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX vendorrefundandcancellationpolicy_vendor_id_idx ON public."VendorRefundAndCancellationPolicy" USING btree ("vendor_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public."VendorRefundAndCancellationPolicy" DROP CONSTRAINT vendorrefundandcancellationpolicy_vendors_fk;
DROP TABLE IF EXISTS public."VendorRefundAndCancellationPolicy";
-- +goose StatementEnd
