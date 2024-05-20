import { z } from 'zod';

export interface Device {
  id: string;
  name: string;
  type: string;
  company_id: string;
  labels: string;
  node_name: string;
  created_at: string;
  updated_at: string;
}

export const createDeviceSchema = z.object({
  name: z
    .string({ message: 'Required' })
    .min(8, { message: 'Must be at least 8 characters' }),
  node_name: z
    .string({ message: 'Required' })
    .min(8, { message: 'Must be at least 8 characters' }),
  type: z.string({ message: 'Required' }),
  labels: z
    .string({ message: 'Required' })
    .includes('=', { message: 'Must be a key value separated by =' }),
});

export type CreateDeviceSchema = z.infer<typeof createDeviceSchema>;
