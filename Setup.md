
1. [Glide]

package management Tool like NPM

[https://glide.readthedocs.io/en/latest/getting-started/](https://glide.readthedocs.io/en/latest/getting-started/)

#### Install
1) mac
brew install glide
2) linux
curl https://glide.sh/get | sh
3) windows
https://github.com/Masterminds/glide/releases
Download : glide-v0.12.3-windows-amd64.zip
extract glide.exe file in zip, and copy it into your
GOPATH/bin. (Please check your GOPATH!)
#### Setup

1) clone project
2) `glide install`

2. [TABLE SCRIPT]

```
DROP TABLE IF EXISTS "USER";

CREATE TABLE public."USER"
(
"ID" serial,
"Firstname" character varying(10) COLLATE pg_catalog."default" NOT NULL,
"Lastname" character varying(10) COLLATE pg_catalog."default" NOT NULL,
"Age" integer NOT NULL,
"Email" character varying(60) COLLATE pg_catalog."default" NOT NULL,
"Address" character varying(100) COLLATE pg_catalog."default" NOT NULL,
"Password" character varying(100) COLLATE pg_catalog."default" NOT NULL,
CONSTRAINT "RECIPE_pkey" PRIMARY KEY ("ID")
)
WITH (
OIDS = FALSE
)
TABLESPACE pg_default;
```