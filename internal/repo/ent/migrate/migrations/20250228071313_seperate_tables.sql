-- Create "cities" table
CREATE TABLE "cities" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "name_en" character varying NOT NULL, "name_cn" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "continents" table
CREATE TABLE "continents" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "name_en" character varying NOT NULL, "name_cn" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "countries" table
CREATE TABLE "countries" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "name_en" character varying NOT NULL, "name_cn" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "states" table
CREATE TABLE "states" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "name_en" character varying NOT NULL, "name_cn" character varying NOT NULL, PRIMARY KEY ("id"));
-- Rename a column from "recommended_duration_seconds" to "recommended_duration_minutes"
ALTER TABLE "points_of_interest" RENAME COLUMN "recommended_duration_seconds" TO "recommended_duration_minutes";
-- Modify "points_of_interest" table
ALTER TABLE "points_of_interest" DROP COLUMN "city", DROP COLUMN "state", DROP COLUMN "country", ADD COLUMN "city_poi" uuid NULL, ADD COLUMN "continent_poi" uuid NULL, ADD COLUMN "country_poi" uuid NULL, ADD COLUMN "state_poi" uuid NULL, ADD CONSTRAINT "points_of_interest_cities_poi" FOREIGN KEY ("city_poi") REFERENCES "cities" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "points_of_interest_continents_poi" FOREIGN KEY ("continent_poi") REFERENCES "continents" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "points_of_interest_countries_poi" FOREIGN KEY ("country_poi") REFERENCES "countries" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "points_of_interest_states_poi" FOREIGN KEY ("state_poi") REFERENCES "states" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
