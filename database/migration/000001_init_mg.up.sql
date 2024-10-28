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

CREATE TABLE "flashcard_progress" (
    "flashcard_id" BIGINT,
    "flashcard_set_id" BIGINT,
    "repetition_number" BIGINT,
    "easiness_factor" DECIMAL(64,2), 
    "interval" BIGINT,
    "last_review" TIMESTAMP DEFAULT NOW(), 
    "due_date" TIMESTAMP DEFAULT NOW() + INTERVAL '1 DAY'  
);

CREATE TABLE "sessions" (
    "session_id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT,
    "flashcard_set_id" BIGINT,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "end_at" TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "session_flashcards" (
    "session_id" BIGINT,
    "flashcard_id" BIGINT,
    "is_reviewed" BOOLEAN DEFAULT FALSE,
    "is_correct" BOOLEAN DEFAULT NULL,
    "review_time" TIMESTAMP
);

ALTER TABLE "session_flashcards" ADD FOREIGN KEY ("session_id") REFERENCES "sessions" ("session_id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "flashcard_progress" ADD FOREIGN KEY ("flashcard_id") REFERENCES "flashcards" ("flashcard_id");

ALTER TABLE "flashcards" ADD FOREIGN KEY ("flashcard_set_id") REFERENCES "flashcard_sets" ("flashcard_set_id");

ALTER TABLE "flashcard_sets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
