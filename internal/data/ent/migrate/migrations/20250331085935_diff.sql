-- Modify "countries" table
ALTER TABLE "countries" ADD COLUMN "file_id" bigint NULL, ADD CONSTRAINT "countries_files_image" FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
