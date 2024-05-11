ALTER TABLE IF EXISTS deployments
ADD COLUMN IF NOT EXISTS company_id UUID;

ALTER TABLE IF EXISTS deployments
ADD CONSTRAINT deployments_company_id_fkey
FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;