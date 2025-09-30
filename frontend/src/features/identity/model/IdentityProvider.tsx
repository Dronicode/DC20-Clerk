import { clearMeProfile } from '@entities/me/model/meProfile';
import React, { createContext, useEffect, useState } from 'react';
import { cacheProfile } from './cacheProfile';
import { onStorageChange } from '@shared/lib/storageListener';
import { STORAGE_KEYS } from '@shared/config/storageKeys';

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
      cacheProfile();
    }

    const unsubscribe = onStorageChange((key, _, newValue) => {
      if (key !== STORAGE_KEYS.accessToken) return;
      if (newValue === null) {
        logout();
        return;
      }
      setIsLoggedIn(true);
      cacheProfile();
    });

    return unsubscribe;
  }, []);

  const logout = () => {
    localStorage.removeItem(STORAGE_KEYS.accessToken);
    setIsLoggedIn(false);
    clearMeProfile();
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, logout }}>
      {children}
    </AuthContext.Provider>
  );
};
