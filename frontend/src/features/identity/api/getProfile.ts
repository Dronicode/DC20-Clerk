import { clearMeProfile } from '@entities/me';
import { STORAGE_KEYS } from '@shared/config';
import type { UserProfile } from '@shared/types';

// Sends a GET request to /identity/profile using the provided token
// Returns the user's profile if successful, or null on failure
export async function getProfile(): Promise<UserProfile | null> {
  try {
    const accessToken = localStorage.getItem(STORAGE_KEYS.accessToken);
    if (!accessToken) {
      console.warn('[PROFILE] No access token available');
      return null;
    }

    console.log('[PROFILE] get /identity/profile using token: %s', accessToken);
    const response = await fetch('/api/identity/profile', {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });
    console.log('[PROFILE] response: %s', response);

    if (response.status === 401) {
      localStorage.removeItem(STORAGE_KEYS.accessToken);
      clearMeProfile();
      return null;
    }

    if (!response.ok) {
      console.error('Response status: %s', response.status);
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
