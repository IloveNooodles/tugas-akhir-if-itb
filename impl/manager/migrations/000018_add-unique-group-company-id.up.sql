ALTER TABLE IF EXISTS groups
ADD CONSTRAINT groups_name_company_id_unique UNIQUE(name, company_id);