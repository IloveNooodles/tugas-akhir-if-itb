ALTER TABLE IF EXISTS groupdevices
ADD CONSTRAINT groupdevices_group_id_device_id_unique UNIQUE(group_id, device_id);