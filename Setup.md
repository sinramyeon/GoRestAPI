
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
CREATE TABLE public."RECIPE"
(
"UNIQUEID" serial,
"NAME" character varying(60) COLLATE pg_catalog."default" NOT NULL,
"PREPTIME" character varying(10) COLLATE pg_catalog."default" NOT NULL,
"DIFFICULTY" integer NOT NULL,
"VEGETARIAN" boolean NOT NULL,
CONSTRAINT "RECIPE_pkey" PRIMARY KEY ("UNIQUEID")
)
WITH (
OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public."RECIPE"
OWNER to hellofresh;
```

3. [API EndPoint]

##### Recipes

| Name | Method | URL | Protected |
| --- | --- | --- | --- |
| List | `GET` | `/recipes` | ✘ |
| List | `GET` | `/recipes?p={p}` | ✘ |
| Create | `POST` | `/recipes` | ✓ |
| Get | `GET` | `/recipes/{id}` | ✘ |
| Update | `PUT/PATCH` | `/recipes/{id}` | ✓ |
| Delete | `DELETE` | `/recipes/{id}` | ✓ |
| Rate | `Post` | `/recipes/{id}/rating` | ✘ |
| Search | `GET` | `/recipes/name/search?q={q}` | ✘ |
| Search | `GET` | `/recipes/time/search?q={q}` | ✘ |
| Search | `GET` | `/recipes/difficulty/search?q={q}` | ✘ |
| Search | `GET` | `/recipes/vegeterain/search?q={q}` | ✘ |
| Search | `GET` | `/recipes/name/search?q={q}&p={p}` | ✘ |
| Search | `GET` | `/recipes/time/search?q={q}&p={p}` | ✘ |
| Search | `GET` | `/recipes/difficulty/search?q={q}&p={p}` | ✘ |
| Search | `GET` | `/recipes/vegeterain/search?q={q}&p={p}` | ✘ |

4. [Security]

- jwt token