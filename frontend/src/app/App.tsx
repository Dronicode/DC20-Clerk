import { RouterProvider } from 'react-router-dom';
import { router } from './providers/router';

export const App = () => {
  return <RouterProvider router={router} />;
};
