import type { UserProfile } from '@shared/types/UserProfile';
import { useEffect, useState } from 'react';
import { getMeProfile, subscribeMeProfile } from '../model/meProfile';

/**
 * React hook to access the current profile reactively.
 */
export const useMeProfile = (): UserProfile | null => {
  const [state, setState] = useState<UserProfile | null>(getMeProfile());

  useEffect(() => {
    const unsubscribe = subscribeMeProfile(setState);
    return () => unsubscribe();
  }, []);

  return state;
};
