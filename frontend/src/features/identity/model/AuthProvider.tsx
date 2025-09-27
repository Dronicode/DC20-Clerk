import { clearMeProfile } from '@entities/me/model/meProfile';
import React, { createContext, useEffect, useState } from 'react';

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
    const access_token = localStorage.getItem('access_token');
    if (!access_token) return;

    setIsLoggedIn(true);
  }, []);

  const logout = () => {
    localStorage.removeItem('access_token');
    setIsLoggedIn(false);
    clearMeProfile();
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, logout }}>
      {children}
    </AuthContext.Provider>
  );
};
