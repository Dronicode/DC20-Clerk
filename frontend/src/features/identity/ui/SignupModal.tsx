import React, { useState } from 'react';
import { setMeProfile } from '@entities/me';
import { getProfile, registerUser } from '../';
import type { UserProfile } from '@shared/types';
import { STORAGE_KEYS } from '@shared/config';

interface SignupModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSignupSuccess: () => void;
  onSwitchToLogin: () => void;
}

export const SignupModal: React.FC<SignupModalProps> = ({
  isOpen,
  onClose,
  onSignupSuccess,
  onSwitchToLogin,
}) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  if (!isOpen) return null;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (password !== confirmPassword) {
      alert('Passwords do not match');
      return;
    }

    try {
      const { access_token } = await registerUser(email, password);
      localStorage.setItem(STORAGE_KEYS.accessToken, access_token);

      const profile: UserProfile | null = await getProfile();
      if (!profile) throw new Error('Failed to load profile');

      setMeProfile(profile);
      onSignupSuccess();
    } catch (err) {
      console.error('[SIGNUP] Failed:', err);
      alert('Signup failed');
    }
  };

  return (
    <div>
      <h2>Sign Up</h2>
      <form onSubmit={handleSubmit}>
        <label>Email</label>
        <input type="email" onChange={(e) => setEmail(e.target.value)} />

        <label>Password</label>
        <input type="password" onChange={(e) => setPassword(e.target.value)} />

        <label>Confirm Password</label>
        <input
          type="password"
          onChange={(e) => setConfirmPassword(e.target.value)}
        />

        <button type="submit">Register</button>
        <button type="button" onClick={onClose}>
          Cancel
        </button>
      </form>
      <p>
        Already have an account?{' '}
        <button type="button" onClick={onSwitchToLogin}>
          Log in here
        </button>
      </p>
    </div>
  );
};
