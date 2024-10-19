CREATE TABLE "users" (
  "user_id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR UNIQUE NOT NULL,
  "password_hash" VARCHAR NOT NULL,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP
);

CREATE TABLE "flashcard_sets" (
  "flashcard_set_id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGINT,
  "title" VARCHAR,
  "description" VARCHAR,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP
);

CREATE TABLE "flashcards" (
  "flashcard_id" BIGSERIAL PRIMARY KEY,
  "flashcard_set_id" BIGINT,
  "term" VARCHAR,
  "definition" VARCHAR,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP
);

ALTER TABLE "flashcards" ADD FOREIGN KEY ("flashcard_set_id") REFERENCES "flashcard_sets" ("flashcard_set_id");

ALTER TABLE "flashcard_sets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
