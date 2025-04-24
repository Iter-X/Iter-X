-- Modify "agent_prompts" table
ALTER TABLE "agent_prompts" DROP COLUMN "rounds", ADD COLUMN "system" character varying NOT NULL, ADD COLUMN "user" character varying NOT NULL;
