CREATE TYPE "user_role" AS ENUM ('librarian', 'member');

CREATE TABLE IF NOT EXISTS "users" (
  "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "username" VARCHAR UNIQUE NOT NULL,
  "password" TEXT NOT NULL,
  "role" user_role NOT NULL DEFAULT 'member',
  "quantity_of_borrow" SMALLINT NOT NULL DEFAULT 3,
  "address" TEXT NOT NULL,
  "image_url" TEXT,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);