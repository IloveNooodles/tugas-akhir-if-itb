import { z } from 'zod';

export interface User {
  ID: string;
  Name: string;
  Email: string;
  Password: string;
  CompanyID: string;
  CreatedAt: string;
  UpdatedAt: string;
}

export const userLoginSchema = z.object({
  email: z.string().email('Invalid email'),
  password: z.string().min(8, 'Must be at least 8 characters'),
});

export type UserLoginSchema = z.infer<typeof userLoginSchema>;
export interface UserLoginResponse {
  data: string;
}
