ALTER TABLE IF EXISTS deployments
ADD COLUMN IF NOT EXISTS target VARCHAR(255) NOT NULL;