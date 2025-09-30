import { setMeProfile } from '@entities/me/model/meProfile';
import { getProfile } from '../api/getProfile';

/**
 * Gets the profile from the backend, passes it to entity Me for caching.
 */
export async function cacheProfile(): Promise<void> {
  const fetched = await getProfile();
  if (!fetched) {
    console.warn('[PROFILE] Failed to fetch the profile from the backend.');
    return;
  }

  setMeProfile(fetched);
}
