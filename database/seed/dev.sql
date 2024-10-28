INSERT INTO "users" ("username", "password_hash", "created_at", "updated_at")
VALUES 
  ('john_doe', 'hashed_password_123', NOW(), NOW()),
  ('jane_smith', 'hashed_password_456', NOW(), NOW());

INSERT INTO "flashcard_sets" ("user_id", "title", "description", "created_at", "updated_at")
VALUES 
  (1, 'Math Basics', 'A set of flashcards for basic math concepts.', NOW(), NOW()),
  (1, 'Spanish Vocabulary', 'Basic Spanish words and phrases.', NOW(), NOW()),
  (2, 'Computer Science 101', 'Introductory computer science topics.', NOW(), NOW());

INSERT INTO "flashcards" ("flashcard_set_id", "term", "definition", "created_at", "updated_at")
VALUES 
  (1, 'Addition', 'The process of adding two or more numbers together.', NOW(), NOW()),
  (1, 'Subtraction', 'The process of taking one number away from another.', NOW(), NOW()),
  (2, 'Hola', 'Hello in Spanish.', NOW(), NOW()),
  (2, 'Gracias', 'Thank you in Spanish.', NOW(), NOW()),
  (3, 'Algorithm', 'A step-by-step procedure for solving a problem.', NOW(), NOW()),
  (3, 'Data Structure', 'A way of organizing data in a computer.', NOW(), NOW());

INSERT INTO "flashcard_progress" ("flashcard_id", "last_review", "repetition_number", "easiness_factor", "interval")
VALUES 
  (1, NOW(), 0, 2.5, 1),
  (2, NOW(), 0, 2.5, 1),
  (3, NOW(), 0, 2.5, 1),
  (4, NOW(), 0, 2.5, 1),
  (5, NOW(), 0, 2.5, 1),
  (6, NOW(), 0, 2.5, 1);

