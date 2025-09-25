export async function getProfile(
  token: string,
): Promise<{ email: string } | null> {
  try {
    console.log('[PROFILE] get using token: ', token);
    const response = await fetch('/identity/profile', {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    console.log('[PROFILE] response: ', response);
    if (!response.ok) {
      throw new Error('Failed to fetch profile');
    }

    const data = await response.json();

    console.log('[PROFILE] Full user profile:', data);
    return { email: data.email };
  } catch (error) {
    console.error('[PROFILE] Failed to fetch profile:', error);
    return null;
  }
}
