CREATE TABLE IF NOT EXISTS deployment_histories (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  device_id UUID NOT NULL,
  image_id UUID NOT NULL,
  deployment_id UUID NOT NULL,
  status VARCHAR(255) NOT NULL DEFAULT "PREPARING",
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (device_id) REFERENCES devices(id),
  FOREIGN KEY (image_id) REFERENCES deployment_repositories(id)
);