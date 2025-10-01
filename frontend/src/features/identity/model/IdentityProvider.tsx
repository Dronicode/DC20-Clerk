import React, { createContext, useEffect, useMemo, useState } from 'react';
import {
  clearMeProfile,
  getMeProfile,
  hydrateMeProfileFromSession,
} from '@entities/me';
import { STORAGE_KEYS } from '@shared/config';
import { onStorageChange } from '@shared/lib';
import { registerLoginStateSetter } from './authSync';
import { cacheProfile } from './cacheProfile';

type AuthContextType = {
  isLoggedIn: boolean;
  logout: () => void;
};

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);

export const IdentityProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const accessToken = localStorage.getItem(STORAGE_KEYS.accessToken);
    if (accessToken) {
      setIsLoggedIn(true);
    }

    registerLoginStateSetter(setIsLoggedIn);

    const unsubscribe = onStorageChange((key, _, newValue) => {
      console.log('Storage Change Signalled');
      if (key !== STORAGE_KEYS.accessToken) return;
      console.log('Storage Change Key true');
      if (newValue === null) {
        console.log('Storage Change newval === null');
        logout();
        return;
      }
      console.log('Storage Change newval not null');
      setIsLoggedIn(true);
      hydrateMeProfileFromSession();
      if (!getMeProfile()) {
        cacheProfile();
      }
    });

    return unsubscribe;
  }, []);

  const logout = () => {
    localStorage.removeItem(STORAGE_KEYS.accessToken);
    setIsLoggedIn(false);
    clearMeProfile();
  };

  const contextValue = useMemo(() => {
    return { isLoggedIn, logout };
  }, [isLoggedIn]);

  return (
    <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
  );
};
