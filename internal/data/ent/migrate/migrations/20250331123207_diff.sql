-- Modify "files" table
ALTER TABLE "files" ALTER COLUMN "size" DROP NOT NULL, DROP COLUMN "url", DROP COLUMN "star";
