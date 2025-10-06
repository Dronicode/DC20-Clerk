export async function loginUser(
  email: string,
  password: string,
): Promise<{ access_token: string }> {
  const response = await fetch('/api/identity/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    throw new Error('Login failed');
  }

  const data = await response.json();
  console.log('[LOGIN] Received access_token:', data.access_token);
  return data;
}
