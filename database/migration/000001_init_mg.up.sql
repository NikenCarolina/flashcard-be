CREATE TABLE "users" (
  "user_id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR UNIQUE NOT NULL DEFAULT '',
  "password_hash" VARCHAR NOT NULL DEFAULT '',
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "flashcard_sets" (
  "flashcard_set_id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGINT,
  "title" VARCHAR DEFAULT '',
  "description" VARCHAR DEFAULT '',
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "flashcards" (
  "flashcard_id" BIGSERIAL PRIMARY KEY,
  "flashcard_set_id" BIGINT,
  "term" VARCHAR DEFAULT '',
  "definition" VARCHAR DEFAULT '',
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP DEFAULT NOW()
);

ALTER TABLE "flashcards" ADD FOREIGN KEY ("flashcard_set_id") REFERENCES "flashcard_sets" ("flashcard_set_id");

ALTER TABLE "flashcard_sets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
