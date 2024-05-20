import { z } from 'zod';

export interface GroupDevice {
  id: string;
  group_id: string;
  device_id: string;
  company_id: string;
  created_at: string;
  updated_at: string;
}

export interface GroupDeviceD {
  id: string;
  group_id: string;
  name: string;
}

export interface GroupDeviceG {
  id: string;
  device_id: string;
  name: string;
  type: string;
  labels: string;
}

export const createGroupDeviceSchema = z.object({
  group_id: z.string().uuid({ message: 'invalid uuid' }),
  device_id: z.string().uuid({ message: 'invalid uuid' }),
});

export type CreateGroupDevice = z.infer<typeof createGroupDeviceSchema>;
