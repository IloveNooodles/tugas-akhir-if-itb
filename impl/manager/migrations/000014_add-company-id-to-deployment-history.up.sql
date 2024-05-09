ALTER TABLE IF EXISTS deployment_histories
ADD COLUMN IF NOT EXISTS company_id UUID;

ALTER TABLE IF EXISTS deployment_histories
ADD CONSTRAINT deployment_histories_company_id_fkey
FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;