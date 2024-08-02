CREATE TABLE IF NOT EXISTS "books" (
  "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  "title" VARCHAR NOT NULL,
  "overview" TEXT NOT NULL,
  "image_url" TEXT NOT NULL,
  "isbn" VARCHAR UNIQUE,
  "writer" VARCHAR NOT NULL,
  "publisher" VARCHAR NOT NULL,
  "published_at" SMALLINT NOT NULL,
  "category" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP 
);