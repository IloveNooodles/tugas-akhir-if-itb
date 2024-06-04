import { z } from 'zod';

export interface RemoteDeploy {
  deployment_ids: string[];
  type: string;
}

export const doRemoteDeploySchema = z.object({
  deployment_ids: z.array(
    z.object({
      name: z.string(),
      value: z.string().uuid(),
    }),
  ),
  type: z.string({ message: 'required' }).default('TARGET'),
  custom: z.object({
    kind: z.string(),
    list_id: z.array(z.string()),
  }),
});

export type DoRemoteDeploySchema = z.infer<typeof doRemoteDeploySchema>;

export const deleteDeploySchema = z.object({
  deployment_ids: z.array(z.object({})),
});

export type DeleteDeploySchema = z.infer<typeof deleteDeploySchema>;
