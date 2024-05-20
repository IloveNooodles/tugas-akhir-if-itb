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

export const createDeploymentSchema = z.object({
  repository_id: z
    .string({ message: 'Required' })
    .uuid({ message: 'Must be valid uuid' }),
  name: z
    .string({ message: 'required' })
    .min(8, { message: 'Must be at least 8 characters' }),
  version: z.string().startsWith('v'),
  target: z
    .string({ message: 'Required' })
    .includes('=', { message: 'Must be a key value separated by =' }),
});

export type CreateDeploymentSchema = z.infer<typeof createDeploymentSchema>;
