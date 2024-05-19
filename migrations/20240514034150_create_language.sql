-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public."Language" (
	"id" serial4 NOT NULL,
	"code" varchar(5) NULL,
	"display_name" varchar(100) NULL,
	CONSTRAINT "LanguagePk" PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public."Language";
-- +goose StatementEnd
