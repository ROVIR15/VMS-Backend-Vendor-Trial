-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vendor_payment_policy_code') THEN
        CREATE TYPE vendor_payment_policy_code AS enum ('after_booking_confirmed', 'before_service_date', 'after_service_date', 'after_service_is_fulfilled');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS public."VendorPaymentPolicyValue" (
    "id" serial PRIMARY KEY,
	"vendor_payment_policy_id" int4 NOT NULL,
	"vendor_payment_policy_code" public.vendor_payment_policy_code NOT NULL,
	"charge_percentage" float4 DEFAULT 0 NOT NULL,
	"days" int2 DEFAULT 0 NULL,
	"hours" int2 DEFAULT 0 NULL,
	"minutes" int2 DEFAULT 0 NULL,
	CONSTRAINT vendorpaymentpolicyvalue_unique UNIQUE ("vendor_payment_policy_id"),
	CONSTRAINT vendorpaymentpolicyvalue_vendorpaymentpolicy_fk FOREIGN KEY ("vendor_payment_policy_id") REFERENCES public."VendorPaymentPolicy"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX vendorpaymentpolicyvalue_vendor_payment_policy_id_idx ON public."VendorPaymentPolicyValue" USING btree ("vendor_payment_policy_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public."VendorPaymentPolicyValue" DROP CONSTRAINT IF EXISTS vendorpaymentpolicyvalue_vendorpaymentpolicy_fk;
DROP TABLE IF EXISTS public."VendorPaymentPolicyValue";
SELECT 'down SQL query';
-- +goose StatementEnd
