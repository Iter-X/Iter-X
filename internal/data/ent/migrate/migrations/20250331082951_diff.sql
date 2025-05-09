-- Create "files" table
CREATE TABLE "files" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "object_key" character varying NOT NULL, "size" bigint NOT NULL, "url" character varying NOT NULL, "star" character varying NOT NULL, "ext" character varying NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "files_users_files" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "file_object_key" to table: "files"
CREATE UNIQUE INDEX "file_object_key" ON "files" ("object_key");
-- Create "poi_files" table
CREATE TABLE "poi_files" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "file_id" bigint NOT NULL, "poi_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "poi_files_files_poi_files" FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "poi_files_points_of_interest_poi_files" FOREIGN KEY ("poi_id") REFERENCES "points_of_interest" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "pointsofinterestfiles_poi_id_file_id" to table: "poi_files"
CREATE UNIQUE INDEX "pointsofinterestfiles_poi_id_file_id" ON "poi_files" ("poi_id", "file_id");
