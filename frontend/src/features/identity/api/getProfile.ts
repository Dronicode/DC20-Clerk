import type { UserProfile } from '@shared/types/UserProfile';

// Sends a GET request to /identity/profile using the provided token
// Returns the user's email if successful, or null on failure
export async function getProfile(): Promise<UserProfile | null> {
  try {
    const accessToken = localStorage.getItem('access_token');
    if (!accessToken) {
      console.warn('[PROFILE] No access token available');
      return null;
    }

    console.log('[PROFILE] get /identity/profile using token: %d', accessToken);
    const response = await fetch('/identity/profile', {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });

    console.log('[PROFILE] response: ', response);

    if (!response.ok) {
      throw new Error('Failed to fetch profile');
    }

    // Parse JSON response body
    const profileData = (await response.json()) as UserProfile;
    console.log('[PROFILE] Full user profile:', profileData);

    return profileData;
  } catch (error) {
    console.error('[PROFILE] Failed to fetch profile:', error);
    return null;
  }
}
