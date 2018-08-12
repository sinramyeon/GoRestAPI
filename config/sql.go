package config

const InitSQL = `
DROP TABLE IF EXISTS "TESTUSER";

CREATE TABLE public."TESTUSER"
(
"ID" serial,
"FIRSTNAME" character varying(10) COLLATE pg_catalog."default" NOT NULL,
"LASTNAME" character varying(10) COLLATE pg_catalog."default" NOT NULL,
"AGE" character varying(10)  NOT NULL,
"EMAIL" character varying(60) COLLATE pg_catalog."default" NOT NULL,
"ADDRESS" character varying(100) COLLATE pg_catalog."default" NOT NULL,
"PASSWORD" character varying(100) COLLATE pg_catalog."default" NOT NULL,
CONSTRAINT "TESTUSER_pkey" PRIMARY KEY ("ID")
)
WITH (
OIDS = FALSE
)
TABLESPACE pg_default;

GRANT ALL PRIVILEGES ON DATABASE "test" to admin;
ALTER USER admin WITH SUPERUSER;
`
