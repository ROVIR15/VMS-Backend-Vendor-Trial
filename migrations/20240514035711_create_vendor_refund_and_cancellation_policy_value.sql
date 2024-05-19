-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public."VendorRefundAndCancellationPolicyValue" (
    "ID" SERIAL PRIMARY KEY,
    "vendor_refund_and_cancellation_policy_id" INT4 NOT NULL,
    "is_full_refund" BOOL DEFAULT TRUE NOT NULL,
    "percentage" FLOAT4 DEFAULT 0 NOT NULL,
    "days_prior_more_than_or_equal" INT4 DEFAULT 0 NOT NULL,
    "days_prior_less_than" INT4 DEFAULT 0 NOT NULL,
    "grace_period" INT4 DEFAULT 0 NOT NULL,
    "order" INT4 NULL,
	CONSTRAINT vracpv_vracp_fk FOREIGN KEY ("vendor_refund_and_cancellation_policy_id") REFERENCES public."VendorRefundAndCancellationPolicy"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX vendorrefundandcancellationpolicyvalue_vendor_refund_and_cancel ON public."VendorRefundAndCancellationPolicyValue" USING btree ("vendor_refund_and_cancellation_policy_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public."VendorRefundAndCancellationPolicyValue" DROP CONSTRAINT vracpv_vracp_fk;
DROP TABLE IF EXISTS public."VendorRefundAndCancellationPolicyValue";
-- +goose StatementEnd