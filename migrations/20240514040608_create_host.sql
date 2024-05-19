-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public."Host" (
    "id" serial PRIMARY KEY,
    "vendor_id" int4 NOT NULL,
    "code" varchar(256) NULL,
    "commission" float4 NULL,
    "is_commission_include_tax" bool DEFAULT TRUE NOT NULL,
    "is_allow_booking" bool DEFAULT TRUE NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NULL,
    "updated_at" timestamp NULL,
    "deleted_at" timestamp NULL,
	CONSTRAINT host_vendors_fk FOREIGN KEY ("vendor_id") REFERENCES public."Vendors"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX host_vendor_id_idx ON public."Host" USING btree ("vendor_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public."Host" DROP CONSTRAINT IF EXISTS host_vendors_fk;
DROP TABLE IF EXISTS public."Host";
-- +goose StatementEnd
