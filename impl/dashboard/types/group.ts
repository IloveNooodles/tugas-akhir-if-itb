import { z } from 'zod';

export interface Group {
  id: string;
  name: string;
  company_id: string;
  created_at: string;
  updated_at: string;
}

export const createGroupSchema = z.object({
  name: z
    .string({ message: 'Required' })
    .min(8, { message: 'Minimal 8 Character' }),
});

export type CreateGroupSchema = z.infer<typeof createGroupSchema>;
