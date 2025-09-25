// features/identity/model/AuthProvider.tsx
import React, { createContext, useContext, useEffect, useState } from 'react';
import { setMe, clearMe } from '@entities/me';
import { getProfile } from '../api/getProfile';

type AuthContextType = {
  isLoggedIn: boolean;
  logout: () => void;
};

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('access_token');
    if (!token) return;

    getProfile(token).then((profile) => {
      if (profile?.email) {
        setIsLoggedIn(true);
        setMe({ accessToken: token, email: profile.email });
      }
    });
  }, []);

  const logout = () => {
    localStorage.removeItem('access_token');
    setIsLoggedIn(false);
    clearMe();
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, logout }}>
      {children}
    </AuthContext.Provider>
  );
};
