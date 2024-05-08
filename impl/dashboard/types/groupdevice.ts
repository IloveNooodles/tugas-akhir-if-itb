import { z } from 'zod';

export interface GroupDevice {}

export const createGroupDeviceSchema = z.object({
  group_id: z.string().uuid({ message: 'invalid uuid' }),
  device_id: z.string().uuid({ message: 'invalid uuid' }),
});

export type CreateGroupDevice = z.infer<typeof createGroupDeviceSchema>;
