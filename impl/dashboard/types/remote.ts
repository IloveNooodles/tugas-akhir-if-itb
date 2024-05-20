import { z } from 'zod';

export interface RemoteDeploy {
  deployment_ids: string[];
}

export const doRemoteDeploySchema = z.object({
  deployment_ids: z.array(z.string()),
});

export type DoRemoteDeploySchema = z.infer<typeof doRemoteDeploySchema>;
