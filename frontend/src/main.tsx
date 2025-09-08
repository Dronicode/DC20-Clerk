import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import './index.css';
import App from './App';

const routes = [{ path: '/', element: <App /> }];

const router = createBrowserRouter(routes, {
  basename: import.meta.env.VITE_BASE_PATH || '/',
});

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
);
