CREATE TABLE IF NOT EXISTS "histories" (
  "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "book_id" UUID NOT NULL,
  "borrowed_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "returned_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "fk_histories_user" FOREIGN KEY ("user_id") REFERENCES "users"("id"),
  CONSTRAINT "fk_histories_book" FOREIGN KEY ("book_id") REFERENCES "books"("id")
);