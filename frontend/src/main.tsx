import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import './index.css';
import { AppLayout } from '@widgets/Navbar'; // exposed via index.ts
import { Home } from '@pages/Home'; // exposed via index.ts
import { CharacterSheetManager } from '@pages/CharacterSheetManager'; // exposed via index.ts

const router = createBrowserRouter([
  {
    path: '/',
    element: <AppLayout />, // shared layout
    children: [
      { index: true, element: <Home /> },
      { path: 'characters', element: <CharacterSheetManager /> },
    ],
  },
]);

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
);
