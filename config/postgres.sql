
DROP DATABASE IF EXISTS "ieltof";

Create Database "ieltof";
\connect ieltof

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "fid" text NOT NULL,
  "name" text,
  "url" text
);

CREATE TABLE "friends" (
  "id" SERIAL PRIMARY KEY,
  "uid" integer REFERENCES "users" (id) ON DELETE CASCADE,
  "fid" integer NOT NULL,
  "apt" BOOLEAN
);

CREATE TABLE "messages" (
  "id" SERIAL PRIMARY KEY,
  "uid" integer REFERENCES "users" (id),
  "did" integer REFERENCES "friends" (id) ON DELETE CASCADE,
  "text" text,
  "tmp" BIGINT
);


