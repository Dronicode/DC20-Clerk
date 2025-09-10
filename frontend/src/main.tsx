import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import './index.css';
import AppLayout from '@layouts/AppLayout';
import Homepage from '@pages/Home';
import CharacterSheetManager from '@pages/CharacterSheetManager';

const router = createBrowserRouter([
  {
    path: '/',
    element: <AppLayout />, // shared layout
    children: [
      { index: true, element: <Homepage /> },
      { path: 'characters', element: <CharacterSheetManager /> },
    ],
  },
]);

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
);
