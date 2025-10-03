export async function registerUser(
  email: string,
  password: string,
): Promise<{ access_token: string }> {
  const res = await fetch('/api/identity/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  });

  if (!res.ok) {
    throw new Error('Signup failed');
  }

  return await res.json(); // expects { access_token }
}
