import { Link } from 'react-router-dom';
import { useAuth, LoginModal, SignupModal } from '@features/identity';
import { useState } from 'react';
import { useMeProfile } from '@entities/me/model/meProfile';

export const Navbar = () => {
  const { isLoggedIn, logout } = useAuth();
  const profile = useMeProfile();

  type ModalType = 'login' | 'signup' | null;
  const [activeModal, setActiveModal] = useState<ModalType>(null);

  return (
    <nav>
      <Link to="/">Home</Link>
      <Link to="/characters">Character Sheet Manager</Link>

      {isLoggedIn && profile ? (
        <>
          <span>{profile.email}</span>
          <button onClick={logout}>Logout</button>
        </>
      ) : (
        <>
          <button onClick={() => setActiveModal('login')}>Login</button>
          <button onClick={() => setActiveModal('signup')}>Sign Up</button>
        </>
      )}
      {activeModal === 'login' && (
        <LoginModal
          isOpen={true}
          onClose={() => setActiveModal(null)}
          onLoginSuccess={() => setActiveModal(null)}
          onSwitchToSignup={() => setActiveModal('signup')}
        />
      )}

      {activeModal === 'signup' && (
        <SignupModal
          isOpen={true}
          onClose={() => setActiveModal(null)}
          onSignupSuccess={() => setActiveModal(null)}
          onSwitchToLogin={() => setActiveModal('login')}
        />
      )}
    </nav>
  );
};
