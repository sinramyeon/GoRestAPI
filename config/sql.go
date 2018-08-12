package config

const InitSQL = `

DROP TABLE IF EXISTS "USER";
DROP TABLE IF EXISTS "USER";

CREATE TABLE public."USER"
(
"ID" serial,
"Firstname" character varying(10) COLLATE pg_catalog."default" NOT NULL,
"Lastname" character varying(10) COLLATE pg_catalog."default" NOT NULL,
"Age" integer NOT NULL,
"Email" character varying(60) COLLATE pg_catalog."default" NOT NULL,
"Address" character varying(100) COLLATE pg_catalog."default" NOT NULL,
CONSTRAINT "RECIPE_pkey" PRIMARY KEY ("ID")
)
WITH (
OIDS = FALSE
)
TABLESPACE pg_default;

`
