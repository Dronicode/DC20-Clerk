import { STORAGE_KEYS } from '@shared/config';
import type { UserProfile } from '@shared/types';

let profile: UserProfile | null = null;
let listeners: ((profile: UserProfile | null) => void)[] = [];

/**
 * Sets the profile and notifies all subscribers.
 */
export const setMeProfile = (data: UserProfile) => {
  profile = data;
  sessionStorage.setItem(STORAGE_KEYS.meProfile, JSON.stringify(data));
  listeners.forEach((fn) => fn(profile));
};

/**
 * Clears the profile and notifies subscribers.
 */
export const clearMeProfile = () => {
  profile = null;
  sessionStorage.removeItem(STORAGE_KEYS.meProfile);
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
  const cached = sessionStorage.getItem(STORAGE_KEYS.meProfile);
  if (!cached) return;

  try {
    profile = JSON.parse(cached) as UserProfile;
    listeners.forEach((fn) => fn(profile));
  } catch (err) {
    console.warn('[meProfile] Failed to parse cached profile:', err);
  }
};
