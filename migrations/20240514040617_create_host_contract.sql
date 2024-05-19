-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public."HostContract" (
    "id" serial PRIMARY KEY,
	"host_partner_id" int4 NOT NULL,
	"url" varchar(255) NULL,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	"updated_at" timestamp NULL,
	"deleted_at" timestamp NULL,
	"start_period" date NOT NULL,
	"end_period" date NOT NULL,
	CONSTRAINT hostcontract_host_fk FOREIGN KEY ("host_partner_id") REFERENCES public."Host"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX hostcontract_host_partner_id_idx ON public."HostContract" USING btree ("host_partner_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public."HostContract" DROP CONSTRAINT IF EXISTS hostcontract_host_fk;
DROP TABLE IF EXISTS public."HostContract";
-- +goose StatementEnd
