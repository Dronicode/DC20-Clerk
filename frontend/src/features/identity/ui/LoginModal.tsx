import React, { useState } from 'react';
import { setMeProfile } from '@entities/me';
import { getProfile, loginUser } from '..';
import type { UserProfile } from '@shared/types';
import { STORAGE_KEYS } from '@shared/config';

interface LoginModalProps {
  isOpen: boolean;
  onClose: () => void;
  onLoginSuccess: () => void;
  onSwitchToSignup: () => void;
}

export const LoginModal: React.FC<LoginModalProps> = ({
  isOpen,
  onClose,
  onLoginSuccess,
  onSwitchToSignup,
}) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  if (!isOpen) return null;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const { access_token } = await loginUser(email, password);
      localStorage.setItem(STORAGE_KEYS.accessToken, access_token);

      const profile: UserProfile | null = await getProfile();
      if (!profile) throw new Error('Failed to load profile');
      setMeProfile(profile);
      onLoginSuccess();
    } catch (err) {
      console.error('[LOGIN] Failed:', err);
      alert('Login failed');
    }
  };

  return (
    <div>
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <label>Email</label>
        <input type="email" onChange={(e) => setEmail(e.target.value)} />

        <label>Password</label>
        <input type="password" onChange={(e) => setPassword(e.target.value)} />

        <button type="submit">Login</button>
        <button type="button" onClick={onClose}>
          Cancel
        </button>
      </form>
      <p>
        Don't have an account?{' '}
        <button type="button" onClick={onSwitchToSignup}>
          Sign up here
        </button>
      </p>
    </div>
  );
};
