ALTER TABLE IF EXISTS deployments
DROP CONSTRAINT IF EXISTS deployments_repository_id_fkey;

ALTER TABLE deployments
ADD CONSTRAINT deployments_repository_id_fkey
FOREIGN KEY (repository_id)
REFERENCES deployment_repositories(id)
ON DELETE CASCADE;
