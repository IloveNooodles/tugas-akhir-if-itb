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
  email: z.string().email('Invalid email'),
  password: z.string().min(8, 'Must be at least 8 characters'),
});

export type UserLoginSchema = z.infer<typeof userLoginSchema>;
export interface UserLoginResponse {
  data: {
    access_token: string;
    refresh_token: string;
  };
}

export interface GetUserResponse {
  data: User;
}

export interface GetAllUsersResponse {
  data: Array<User>;
}
