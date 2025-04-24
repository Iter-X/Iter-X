-- Modify "daily_itineraries" table
ALTER TABLE "daily_itineraries" ADD COLUMN "order" smallint NOT NULL;
-- Modify "trip_collaborators" table
ALTER TABLE "trip_collaborators" ALTER COLUMN "status" SET DEFAULT 'Invited';
