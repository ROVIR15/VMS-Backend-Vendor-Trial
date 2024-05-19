-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public."VendorPaymentPolicy" (
    "id" serial PRIMARY KEY,
	"vendor_id" int4 NOT NULL,
	"is_collect_after_fulfillment" bool DEFAULT true NOT NULL,
	CONSTRAINT vendorpaymentpolicy_unique UNIQUE ("vendor_id"),
	CONSTRAINT vendorpaymentpolicy_vendors_fk FOREIGN KEY ("vendor_id") REFERENCES public."Vendors"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX vendorpaymentpolicy_vendor_id_idx ON public."VendorPaymentPolicy" USING btree ("vendor_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public."VendorPaymentPolicy" DROP CONSTRAINT IF EXISTS vendorpaymentpolicy_vendors_fk;
DROP TABLE IF EXISTS public."VendorPaymentPolicy";
-- +goose StatementEnd
