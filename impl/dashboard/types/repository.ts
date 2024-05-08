import { z } from 'zod';

export interface Repository {
  id: string;
  name: string;
  description: string;
  image: string;
  created_at: string;
  updated_at: string;
}

export const createRepositorySchema = z.object({
  name: z.string().min(8, { message: 'Must be 8 charaters or greater' }),
  description: z.string().min(8, { message: 'Must be 8 charaters or greater' }),
  image: z.string({ message: 'required' }),
});

export type CreateRepositorySchema = z.infer<typeof createRepositorySchema>;
