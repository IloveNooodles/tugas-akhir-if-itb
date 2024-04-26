import type { UserLoginResponse, UserLoginSchema } from '~/types/user';

export async function login(body: UserLoginSchema) {
  const fetch = useAPI()
  const response = fetch<UserLoginResponse>('/users/login', {
    method: 'POST',
    body: body,
  });

  return response;
}
