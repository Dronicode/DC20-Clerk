import {
  getMeProfile,
  hydrateMeProfileFromSession,
  useMeProfile,
} from '@entities/me';
import { cacheProfile } from '@features/identity';
import { useEffect, useState } from 'react';

/**
 * Hydrates meProfile if not already cached.
 * Can be used lazily in components that need profile access.
 */
export function useHydrateMeProfile() {
  const profile = useMeProfile();
  const [hydrated, setHydrated] = useState(!!getMeProfile());

  useEffect(() => {
    if (hydrated) return;

    hydrateMeProfileFromSession();

    if (!getMeProfile()) {
      cacheProfile();
    }

    setHydrated(true);
  }, [hydrated]);

  return { profile, loading: !hydrated };
}
