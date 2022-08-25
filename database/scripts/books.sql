CREATE table "book" (
	"id" SERIAL,
	"author" varchar,
	"title" VARCHAR not null,
	"description" VARCHAR,
	"rating" int,
	"created_at" timestamptz NOT NULL DEFAULT (now()),
	"modified_at" timestamptz NOT NULL DEFAULT (now()),
	"disabled" BOOLEAN not null default FALSE,
	"disabled_at" timestamptz,
	PRIMARY KEY("id")
);