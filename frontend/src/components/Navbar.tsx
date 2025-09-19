import { Link } from 'react-router-dom';
import LoginModal from '@features/identity/components/LoginModal';
import SignupModal from '@features/identity/components/SignupModal';
import React, { useState, useEffect } from 'react';

const Navbar: React.FC = () => {
  const [activeModal, setActiveModal] = useState<'login' | 'signup' | null>(
    null,
  );
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  // Check for token on mount
  useEffect(() => {
    const token = localStorage.getItem('access_token');
    setIsLoggedIn(!!token);
  }, []);

  // Logout handler
  const handleLogout = () => {
    localStorage.removeItem('access_token');
    setIsLoggedIn(false);
  };

  return (
    <>
      <nav>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/characters">Character Sheets</Link>
          </li>
          <li>
            <Link to="/">Other stuff coming later</Link>
          </li>
        </ul>
        {isLoggedIn ? (
          <button onClick={handleLogout}>Logout</button>
        ) : (
          <>
            <button onClick={() => setActiveModal('login')}>Login</button>
            <button onClick={() => setActiveModal('signup')}>Register</button>
          </>
        )}
      </nav>
      <LoginModal
        isOpen={activeModal === 'login'}
        onClose={() => setActiveModal(null)}
      />
      <SignupModal
        isOpen={activeModal === 'signup'}
        onClose={() => setActiveModal(null)}
      />
    </>
  );
};

export default Navbar;
