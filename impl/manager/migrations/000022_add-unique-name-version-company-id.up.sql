ALTER TABLE IF EXISTS deployments
ADD CONSTRAINT deployments_company_id_name_version_unique UNIQUE(name, version, company_id);