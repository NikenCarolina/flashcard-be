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

INSERT INTO "flashcard_progress" ("flashcard_id", "flashcard_set_id", "repetition_number", "easiness_factor", "interval", "last_review", "due_date")
VALUES 
  (1, 1, 0, 2.5, 1, NOW() - INTERVAL '2 DAYS', NOW() + INTERVAL '1 DAY'), 
  (2, 1, 1, 2.5, 3, NOW() - INTERVAL '1 DAY', NOW() + INTERVAL '2 DAYS'), 
  (3, 2, 0, 2.5, 1, NOW() - INTERVAL '4 DAYS', NOW() + INTERVAL '1 DAY'), 
  (4, 2, 1, 2.5, 2, NOW() - INTERVAL '2 DAYS', NOW() + INTERVAL '3 DAYS'), 
  (5, 3, 0, 2.5, 1, NOW() - INTERVAL '3 DAYS', NOW() + INTERVAL '1 DAY'), 
  (6, 3, 2, 2.5, 5, NOW() - INTERVAL '1 DAY', NOW() + INTERVAL '6 DAYS'); 
