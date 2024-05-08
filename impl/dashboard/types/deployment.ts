import { z } from 'zod';

export interface Deployment {
  id: string;
  repository_id: string;
  name: string;
  version: string;
  target: string;
  created_at: string;
  updated_at: string;
}

// TODO: benerin definisi target -> bisa deploy ke group / deviceId
// harus ditambahin logicnya di backend
export const createDeploymentSchema = z.object({
  repository_id: z.string().uuid({ message: 'Must be valid uuid' }),
  name: z
    .string()
    .min(8, { message: 'Must be higher or equal to 8 characters' }),
  version: z.string(),
  target: z.string(),
});

export type CreateDeploymentSchema = z.infer<typeof createDeploymentSchema>;
