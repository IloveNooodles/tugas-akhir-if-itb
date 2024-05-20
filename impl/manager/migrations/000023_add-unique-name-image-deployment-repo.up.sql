ALTER TABLE IF EXISTS deployment_repositories
ADD CONSTRAINT deployments_repositories_name_image_unique UNIQUE(name, image);