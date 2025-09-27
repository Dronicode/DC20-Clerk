import type { UserProfile } from '@shared/types/UserProfile';
import { useEffect, useState } from 'react';

const SESSION_KEY = 'identity.profile';

let profile: UserProfile | null = null;
let listeners: ((profile: UserProfile | null) => void)[] = [];

/**
 * Sets the profile and notifies all subscribers.
 */
export const setMeProfile = (data: UserProfile) => {
  profile = data;
  sessionStorage.setItem(SESSION_KEY, JSON.stringify(data));
  listeners.forEach((fn) => fn(profile));
};

/**
 * Clears the profile and notifies subscribers.
 */
export const clearMeProfile = () => {
  profile = null;
  sessionStorage.removeItem(SESSION_KEY);
  listeners.forEach((fn) => fn(profile));
};

/**
 * Returns the current profile synchronously.
 */
export const getMeProfile = (): UserProfile | null => profile;

/**
 * Subscribes to profile changes and returns an unsubscribe function.
 */
export const subscribeMeProfile = (
  fn: (profile: UserProfile | null) => void,
) => {
  listeners.push(fn);
  return () => {
    listeners = listeners.filter((l) => l !== fn);
  };
};

/**
 * Hydrates profile from sessionStorage if available.
 */
export const hydrateMeProfileFromSession = () => {
  const cached = sessionStorage.getItem(SESSION_KEY);
  if (!cached) return;

  try {
    profile = JSON.parse(cached) as UserProfile;
    listeners.forEach((fn) => fn(profile));
  } catch (err) {
    console.warn('[meProfile] Failed to parse cached profile:', err);
  }
};

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
