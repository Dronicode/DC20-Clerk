import React from 'react';

interface SignupModalProps {
  isOpen: boolean;
  onClose: () => void;
}

const SignupModal: React.FC<SignupModalProps> = ({ isOpen, onClose }) => {
  if (!isOpen) return null;

  return (
    <div className="modal-backdrop">
      <div className="modal-content">
        <h2>Sign Up</h2>
        <form>
          <label>Email</label>
          <input type="email" name="email" />

          <label>Password</label>
          <input type="password" name="password" />

          <label>Confirm Password</label>
          <input type="password" name="confirmPassword" />

          <button type="submit">Register</button>
          <button type="button" onClick={onClose}>
            Cancel
          </button>
        </form>
      </div>
    </div>
  );
};

export default SignupModal;
