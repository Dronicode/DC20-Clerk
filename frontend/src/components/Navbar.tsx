import { Link } from 'react-router-dom';
import LoginModal from '@features/identity/components/LoginModal';
import React, { useState } from 'react';

const Navbar: React.FC = () => {
  const [showLogin, setShowLogin] = useState(false);

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
        </ul>
        <button onClick={() => setShowLogin(true)}>Login</button>
      </nav>
      <LoginModal isOpen={showLogin} onClose={() => setShowLogin(false)} />
    </>
  );
};

export default Navbar;
