-- +goose Up
-- +goose StatementBegin
CREATE TABLE public."VendorAddressTranslation" (
	"vendor_address_id" int4 NOT NULL,
	"language_id" int4 NOT NULL,
	"address" text NULL,
	CONSTRAINT vendoraddresstranslation_pk PRIMARY KEY ("vendor_address_id", "language_id")
);
CREATE INDEX vendoraddresstranslation_company_address_id_idx ON public."VendorAddressTranslation" USING btree ("vendor_address_id", "language_id");

-- public."VendorAddressTranslation" foreign keys
ALTER TABLE public."VendorAddressTranslation" ADD CONSTRAINT vendoraddresstranslation_language_fk FOREIGN KEY ("language_id") REFERENCES public."Language"("id") ON DELETE SET NULL ON UPDATE SET NULL;
ALTER TABLE public."VendorAddressTranslation" ADD CONSTRAINT vendor_address_translation_vendor_address_fk FOREIGN KEY ("vendor_address_id") REFERENCES public."VendorAddress"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public."VendorAddressTranslation" DROP CONSTRAINT vendoraddresstranslation_pk;
DROP TABLE public."VendorAddressTranslation";
-- +goose StatementEnd
