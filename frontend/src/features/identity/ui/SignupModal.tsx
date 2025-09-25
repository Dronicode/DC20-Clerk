import React, { useState } from 'react';
import { setMe } from '@entities/me';
import { getProfile, registerUser } from '../api';

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

      localStorage.setItem('access_token', access_token);

      const profile = await getProfile(access_token);
      if (profile?.email) {
        setMe({ accessToken: access_token, email: profile.email });
        onSignupSuccess();
      } else {
        throw new Error('Failed to load profile');
      }
    } catch (err) {
      console.error(err);
      alert('Signup failed');
    }
  };

  return (
    <div>
      <h2>Sign Up</h2>
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

        <label>Confirm Password</label>
        <input
          type="password"
          name="confirmPassword"
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
