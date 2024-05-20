ALTER TABLE IF EXISTS devices
DROP CONSTRAINT IF EXISTS devices_company_id_node_name_unique;

ALTER TABLE IF EXISTS devices
ADD CONSTRAINT devices_name_company_id_unique UNIQUE(name, company_id);
