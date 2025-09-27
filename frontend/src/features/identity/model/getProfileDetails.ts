import { getMeProfile, setMeProfile } from '@entities/me/model/meProfile';
import { type UserProfile } from '@shared/types/UserProfile';
import { getProfile } from '../api/getProfile';

/**
 * Returns the current profile, either from memory or by fetching from backend.
 */
export async function getProfileDetails(): Promise<UserProfile | null> {
  const cached = getMeProfile();
  if (cached) {
    console.log('[PROFILE] Returning in-memory profile');
    return cached;
  }

  const fetched = await getProfile();
  if (fetched) {
    setMeProfile(fetched);
    return fetched;
  }

  return null;
}
