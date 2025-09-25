import { Outlet } from 'react-router-dom';
import { Navbar } from '@widgets/Navbar';

export const AppLayout = () => {
  return (
    <>
      <Navbar />
      <main>
        <Outlet />
      </main>
    </>
  );
};
