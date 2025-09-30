import React, { createContext, useEffect, useMemo, useState } from 'react';
import { clearMeProfile } from '@entities/me';
import { onStorageChange } from '@shared/lib';
import { registerLoginStateSetter } from './authSync';
import { STORAGE_KEYS } from '@shared/config';

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
      if (key !== STORAGE_KEYS.accessToken) return;
      if (newValue === null) {
        logout();
        return;
      }
      setIsLoggedIn(true);
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
