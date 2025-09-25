import React, { useState } from 'react';
import { setMe } from '@entities/me';
import { getProfile, loginUser } from '../api';

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
      localStorage.setItem('access_token', access_token);

      const profile = await getProfile(access_token);
      if (profile?.email) {
        setMe({ accessToken: access_token, email: profile.email });
        onLoginSuccess();
      } else {
        throw new Error('Failed to load profile');
      }
    } catch (err) {
      console.error(err);
      alert('Login failed');
    }
  };

  return (
    <div>
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <label>Email</label>
        <input
          type="email"
          name="email"
          onChange={(e) => setEmail(e.target.value)}
        />

        <label>Password</label>
        <input
          type="password"
          name="password"
          onChange={(e) => setPassword(e.target.value)}
        />

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
