import { z } from 'zod';

export interface Device {
  id: string;
  name: string;
  type: string;
  company_id: string;
  attributes: string;
  node_name: string;
  created_at: string;
  updated_at: string;
}

// TODO add backend validation minimal name 8
// TODO attribute don't know what for
// TODO type yang bakal jadi label?
// TODO attributes kayaknya extra aja
// TODO node name gaboleh duplicate
// type=attribute? attribute -> x=y
export const createDeviceSchema = z.object({
  name: z.string().min(8, { message: 'Must be at least 8 characters' }),
  node_name: z.string().min(8, { message: 'Must be at least 8 characters' }),
  type: z.string({ message: 'Required' }),
  attributes: z.string({ message: 'Required' }),
});

export type CreateDeviceSchema = z.infer<typeof createDeviceSchema>;
