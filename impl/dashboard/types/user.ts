import { z } from 'zod';

export interface User {
  id: string;
  name: string;
  email: string;
  company_id: string;
  created_at: string;
  updated_at: string;
}

export const userLoginSchema = z.object({
  email: z.string({ message: 'Required' }).email('Invalid email'),
  password: z
    .string({ message: 'Required' })
    .min(8, 'Must be at least 8 characters')
    .max(72, 'Must be shorter than 72 characters'),
});

export type UserLoginSchema = z.infer<typeof userLoginSchema>;
export interface UserLoginResponse {
  data: {
    access_token: string;
    refresh_token: string;
  };
}
