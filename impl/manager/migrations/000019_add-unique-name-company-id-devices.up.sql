ALTER TABLE IF EXISTS devices
ADD CONSTRAINT devices_name_company_id_unique UNIQUE(name, company_id);